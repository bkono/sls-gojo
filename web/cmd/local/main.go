package main

import (
		"log"
	"net/http"

	"github.com/bkono/sls-gojo/web"
)

func main() {
	router := web.NewHandler()
	log.Println("starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

