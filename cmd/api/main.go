package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0xmlx/nigeria-uni/internal/router"
)

// getting our environment variable...
var port = os.Getenv("PORT")

func main() {

	log.Println("server running at port", port)

	route := router.NEW()

	if err := http.ListenAndServe(":"+port, route); err != nil {
		log.Fatal(err)
	}
}
