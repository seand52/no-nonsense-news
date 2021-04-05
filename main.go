package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

// detail structs
type Response struct {
	Response Content `json:"response"`
}

type Content struct {
	Content NewsData `json:"content"`
}

type NewsData struct {
	WebTitle           string     `json:"webTitle"`
	WebPublicationDate string     `json:"webPublicationDate"`
	WebUrl             string     `json:"webUrl"`
	Blocks             NewsBlocks `json:"blocks"`
}

type NewsBlocks struct {
	Main MainData   `json:"main"`
	Body []BodyData `json:"body"`
}

type MainData struct {
	BodyHtml string `json:"bodyHtml"`
}

type BodyData struct {
	BodyDataHtml string `json:"bodyHtml"`
}

//overview structs
type ResponseOverview struct {
	Response Results `json:"response"`
}

type Results struct {
	Result []Result `json:"results"`
}

type Result struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	WebPublicationDate string `json:"webPublicationDate"`
	WebTitle           string `json:"webTitle"`
	ApiUrl             string `json:"apiUrl"`
}

type FilteredResult struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	WebPublicationDate string `json:"webPublicationDate"`
	WebTitle           string `json:"webTitle"`
	Slug               string `json:"slug"`
}

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		//hondle the error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	re := regexp.MustCompile(`([^\/]+$)`)
	var news ResponseOverview
	json.Unmarshal(byteValue, &news)
	var filteredNews []FilteredResult
	for _, v := range news.Response.Result {
		if v.Type != "liveblog" {
			id := string(re.Find([]byte(v.Id)))
			var test FilteredResult
			test.Id = v.Id
			test.Type = v.Type
			test.WebPublicationDate = v.WebPublicationDate
			test.WebTitle = v.WebTitle
			test.Slug = id
			filteredNews = append(filteredNews, test)
		}
	}
	t, err := template.ParseFiles("edit.html")
	if err != nil {
		//hondle the error
		fmt.Println(err)
	}
	err = t.Execute(w, filteredNews)
	if err != nil {
		//hondle the error
		fmt.Println(err)
	}
	resp, err := http.Get("https://content.guardianapis.com/search?api-key=a3136ec7-05ca-42c8-b0ac-60f9eae61e85&page-size=3")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	var apiResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&apiResult)
	fmt.Println(apiResult)

	file, _ := json.MarshalIndent(apiResult, "", " ")
	_ = ioutil.WriteFile("data/test2.json", file, 0644)
}

func fetchNewsdetail() {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		//hondle the error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var news ResponseOverview
	json.Unmarshal(byteValue, &news)
	for _, v := range news.Response.Result {
		if v.Type != "liveblog" {
			apiUrl := v.ApiUrl + "?api-key=a3136ec7-05ca-42c8-b0ac-60f9eae61e85&show-blocks=all"
			fmt.Println("============================== before request")
			resp, err := http.Get(apiUrl)
			if err != nil {
				// handle error
			}
			defer resp.Body.Close()
			// body, err := ioutil.ReadAll(resp.Body)
			var apiResult map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&apiResult)
			file, _ := json.MarshalIndent(apiResult, "", " ")
			re := regexp.MustCompile(`([^\/]+$)`)
			slug := string(re.Find([]byte(v.Id)))
			err = ioutil.WriteFile("data/"+slug+".json", file, 0644)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("============================== before sleep")
      time.Sleep(2*time.Second)
			fmt.Println("============================== after sleep")

		}
	}

}

func renderNewsDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fetchNewsdetail()
	jsonFile, err := os.Open("response.json")
	if err != nil {
		//hondle the error
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		//hondle the error
	}
	var newsDetail Response
	json.Unmarshal(byteValue, &newsDetail)
	fmt.Println(newsDetail)

	articleId := p.ByName("articleId")
	fmt.Println(articleId)
	t, err := template.ParseFiles("newsDetail.html")
	if err != nil {
		//hondle the error
	}
	t.Execute(w, newsDetail)
}

func main() {
	router := httprouter.New()
	router.GET("/", handler)
	router.GET("/article/:articleId", renderNewsDetail)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
