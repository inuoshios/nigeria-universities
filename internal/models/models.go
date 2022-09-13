package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	Universities struct {
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		WebsiteLink  string `json:"website_link"`
	}

	UniversitiesReplica struct {
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		WebsiteLink  string `json:"website_link"`
	}
)

// CovertJSONToStruct reads the JSON file and converts it to a struct...
func CovertJSONToStruct() []Universities {
	uni, err := ioutil.ReadFile("internal/json/uni.json")
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	var data []Universities

	if err = json.Unmarshal(uni, &data); err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	return data
}
