package helpers

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

func WriteDataToJson(data map[string]interface{}, path string) error {
	file, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile(path, file, 0644)
	return err
}

func GetArticleSlug(id string) string {
	re := regexp.MustCompile(`([^\/]+$)`)
	return string(re.Find([]byte(id)))
}
