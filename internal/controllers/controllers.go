package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ixxiv/nigeria-uni/internal/models"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	welcome := map[string]any{
		"status": 200,
		"title":  "List of all univeristies in Nigeria—still in progress.",
		"msg":    "PRs are allowed. 🪢",
	}

	// converts map into JSON format
	convToJson, err := json.MarshalIndent(welcome, "", "    ")
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	w.Write(convToJson)
}

func GetAllUniveristy(w http.ResponseWriter, r *http.Request) {
	var universities []models.Universities

	for _, data := range models.CovertJSONToStruct() {
		universities = append(universities, models.Universities{
			ID:           data.ID,
			Name:         data.Name,
			Abbreviation: data.Abbreviation,
			WebsiteLink:  data.WebsiteLink,
		})
	}

	if err := json.NewEncoder(w).Encode(universities); err != nil {
		log.Fatal(err)
	}
}

func GetSpecificUniversity(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	for _, data := range models.CovertJSONToStruct() {

		if strings.EqualFold(data.Abbreviation, param["abbreviation"]) {
			if err := json.NewEncoder(w).Encode(data); err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	w.Write([]byte("sorry, this resource is not available..."))
}
