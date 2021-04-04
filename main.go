package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

//detail structs
// type Response struct {
//   Response Content `json:"response"`
// }

// type Content struct {
//   Content NewsData `json:"content"`
// }

// type NewsData struct {
// 	WebTitle string `json:"webTitle"`
// }

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
}


func handler(w http.ResponseWriter, r *http.Request) {
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
	err = t.Execute(w, filteredNews)
	if err != nil {
		//hondle the error
		fmt.Println(err)
	}

}
func main() {
	http.HandleFunc("/", handler)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
