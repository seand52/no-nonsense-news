package routes

import (
	"encoding/json"
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
	resp, err := http.Get("https://content.guardianapis.com/search?api-key=a3136ec7-05ca-42c8-b0ac-60f9eae61e85&page-size=50")
	if err != nil {
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
	fetchNewsdetail()

}

func fetchNewsdetail() {
	jsonFile, err := os.Open("data/overview.json")
	if err != nil {
		//hondle the error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var news ResponseOverview
	json.Unmarshal(byteValue, &news)
	for _, v := range news.Response.Result {
		if v.Type != "liveblog" {
			apiUrl := v.ApiUrl + "?api-key=a3136ec7-05ca-42c8-b0ac-60f9eae61e85&show-blocks=all"
			resp, err := http.Get(apiUrl)
			if err != nil {
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
