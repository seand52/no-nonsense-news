package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"no-nonsense-news/helpers"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)
func FetchNews(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  os.RemoveAll("data")
  os.Mkdir("data", os.ModePerm)
  apiKey := os.Getenv("AUTH_KEY")
	resp, err := http.Get("https://content.guardianapis.com/search?api-key=" + apiKey + "&page-size=50")
	if err != nil {
    fmt.Println("fetch error")
		// handle error
	}
	defer resp.Body.Close()
	var apiResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&apiResult)
	error := helpers.WriteDataToJson(apiResult, "data/overview.json")
	if error != nil {
		// handle error
	}
	time.Sleep(5 * time.Second)
	fetchNewsdetail(apiKey)

}

func fetchNewsdetail(apiKey string) {
	jsonFile, err := os.Open("data/overview.json")
	if err != nil {
		//hondle the error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var news ResponseOverview
	json.Unmarshal(byteValue, &news)
	for _, v := range news.Response.Result {
		if v.Type != "liveblog" {
			apiUrl := v.ApiUrl + "?api-key=" + apiKey + "&show-blocks=all"
			resp, err := http.Get(apiUrl)
			if err != nil {
    fmt.Println("fetch error")
				// handle error
			}
			defer resp.Body.Close()
			var apiResult map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&apiResult)
      slug := helpers.GetArticleSlug(v.Id)
			error := helpers.WriteDataToJson(apiResult, "data/"+slug+".json")
			if error != nil {
				// handle error
			}
			time.Sleep(2 * time.Second)

		}
	}

}
