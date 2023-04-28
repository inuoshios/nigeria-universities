package main

import (
	"log"
	"net/http"

	"github.com/inuoshios/nigeria-uni/router"
)

func main() {
	port := "8000"

	log.Println("server running at port", port)

	route := router.NEW()

	if err := http.ListenAndServe(":"+port, route); err != nil {
		log.Fatal(err)
	}
}
