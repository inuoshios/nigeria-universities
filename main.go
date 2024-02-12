package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

type Universities struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	WebsiteLink  string `json:"website_link"`
}

// CovertJSONToStruct reads the JSON file and converts it to a struct...
func CovertJSONToStruct() []Universities {
	uni, err := os.ReadFile("json/uni.json")
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	var data []Universities

	if err = json.Unmarshal(uni, &data); err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	return data
}

func AddContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	welcome := map[string]any{
		"status": 200,
		"title":  "List of all univeristies in Nigeria—still in progress.",
		"msg":    "PRs are allowed. 🪢",
	}
	json.NewEncoder(w).Encode(welcome)
}

func GetAllUniveristy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var universities []Universities
	for index, data := range CovertJSONToStruct() {
		universities = append(universities, Universities{
			ID:           uint(index + 1),
			Name:         data.Name,
			Abbreviation: data.Abbreviation,
			WebsiteLink:  data.WebsiteLink,
		})
	}

	if err := json.NewEncoder(w).Encode(universities); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "An error occurred",
		})
		return
	}
}

func GetSpecificUniversity(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("abbreviation")

	for index, data := range CovertJSONToStruct() {
		data.ID = uint(index + 1)
		if strings.EqualFold(data.Abbreviation, param) {
			if err := json.NewEncoder(w).Encode(data); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "An error occurred",
				})
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Error finding a university with that abbreviation",
	})
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /api", AddContentType(HandlerFunc(Welcome)))
	mux.Handle("GET /api/v1", AddContentType(HandlerFunc(GetAllUniveristy)))
	mux.Handle("GET /api/v1/{abbreviation}", AddContentType(HandlerFunc(GetSpecificUniversity)))

	return mux
}

func main() {
	port := "8000"

	log.Println("server running at port", port)

	if err := http.ListenAndServe(":"+port, routes()); err != nil {
		log.Fatal(err)
	}
}
