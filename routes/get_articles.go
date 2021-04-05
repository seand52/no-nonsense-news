package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"no-nonsense-news/helpers"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

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

type ResultWithSlug struct {
  Result
	Slug               string `json:"slug"`
}

func GetArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonFile, err := os.Open("data/overview.json")
	if err != nil {
		//hondle the error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var news ResponseOverview
	json.Unmarshal(byteValue, &news)
	var filteredNews []ResultWithSlug
	for _, v := range news.Response.Result {
		if v.Type != "liveblog" {
      slug := helpers.GetArticleSlug(v.Id)
      filteredResult := ResultWithSlug{Result: v, Slug: slug}
			filteredNews = append(filteredNews, filteredResult)
		}
	}
	t, err := template.ParseFiles("views/overview.html")
	if err != nil {
		//hondle the error
	}
	err = t.Execute(w, filteredNews)
	if err != nil {
		//hondle the error
	}
}
