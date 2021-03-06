package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"no-nonsense-news/helpers"
	"os"
	"path/filepath"
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
	Slug string `json:"slug"`
}

type RenderData struct {
	Results []ResultWithSlug
	Date    string
}

func GetArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonPath, err := filepath.Abs("./data/overview.json")
	jsonFile, err := os.Open(jsonPath)
	currentDate := helpers.GetFormattedDate()
	if err != nil {
		//hondle the error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var news ResponseOverview
	json.Unmarshal(byteValue, &news)
	var filteredNews []ResultWithSlug
	for _, v := range news.Response.Result {
		if v.Type == "article" {
			slug := helpers.GetArticleSlug(v.Id)
			filteredResult := ResultWithSlug{Result: v, Slug: slug}
			filteredNews = append(filteredNews, filteredResult)
		}
	}
	htmlPath, err := filepath.Abs("./views/overview.html")
	t, err := template.ParseFiles(htmlPath)
	if err != nil {
		//hondle the error
	}
  data := RenderData{Results: filteredNews, Date: currentDate}
	err = t.Execute(w, data)
	if err != nil {
		//hondle the error
	}
}
