package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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
	jsonFile, err := os.Open("data/" + articleId + ".json")
	if err != nil {
		//hondle the error
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		//hondle the error
	}
	var newsDetail Response
	json.Unmarshal(byteValue, &newsDetail)
	t, err := template.ParseFiles("views/news_detail.html")
	if err != nil {
		//hondle the error
	}
	t.Execute(w, newsDetail)
}

