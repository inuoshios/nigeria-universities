package main

import (
	"log"
	"net/http"
	"os"

	"github.com/inuoshios/nigeria-uni/router"
	"github.com/joho/godotenv"
)

func main() {
	// load the env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file %s", err)
	}

	// getting our env
	port := os.Getenv("PORT")

	log.Println("server running at port", port)

	route := router.NEW()

	if err := http.ListenAndServe(":"+port, route); err != nil {
		log.Fatal(err)
	}
}
