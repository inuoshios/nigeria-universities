package models

import (
	"encoding/json"
	"log"
	"os"
)

type Universities struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	WebsiteLink  string `json:"website_link"`
}

// CovertJSONToStruct reads the JSON file and converts it to a struct...
func CovertJSONToStruct() []Universities {
	uni, err := os.ReadFile("internal/json/uni.json")
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	var data []Universities

	if err = json.Unmarshal(uni, &data); err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	return data
}
