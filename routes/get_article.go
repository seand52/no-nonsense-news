package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

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

func GetArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	articleId := p.ByName("articleId")
	jsonPath, err := filepath.Abs("./data/" + articleId + ".json")
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		//hondle the error
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		//hondle the error
	}
	var newsDetail Response
	json.Unmarshal(byteValue, &newsDetail)
	htmlPath, err := filepath.Abs("./views/news_detail.html")
	t, err := template.ParseFiles(htmlPath)
	if err != nil {
		//hondle the error
	}
	t.Execute(w, newsDetail)
}
