package main

import (
	"log"

	"github.com/apex/gateway"
	"github.com/bkono/sls-gojo/web"
)

func main() {
	router := web.NewHandler()
	log.Fatal(gateway.ListenAndServe(":3000", router))
}
