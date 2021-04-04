package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
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
	Slug           string `json:"slug"`
}

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonFile, err := os.Open("overview.json")
	if err != nil {
		//hondle the error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var news ResponseOverview
	json.Unmarshal(byteValue, &news)
	var filteredNews []Result
	for _, v := range news.Response.Result {
		if v.Type != "liveblog" {
			filteredNews = append(filteredNews, v)
		}
	}
	t, err := template.ParseFiles("edit.html")
	if err != nil {
		//hondle the error
		fmt.Println(err)
	}
  fmt.Println(filteredNews)
	err = t.Execute(w, filteredNews)
	if err != nil {
		//hondle the error
		fmt.Println(err)
	}

}

func renderNewsDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	jsonFile, _ := os.Open("response.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var newsDetail Response
	json.Unmarshal(byteValue, &newsDetail)
	fmt.Println(newsDetail)

	articleId := p.ByName("articleId")
	fmt.Println(articleId)
	t, _ := template.ParseFiles("newsDetail.html")
	t.Execute(w, newsDetail)

}
func main() {
	router := httprouter.New()
	router.GET("/", handler)
	router.GET("/article/:articleId", renderNewsDetail)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  router.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
