package main

import (
	"log"

	"github.com/apex/gateway"
	"github.com/bkono/sls-gojo/web"
)

func main() {
	s := web.NewServer()
	log.Println("starting on :3000")
	log.Fatal(gateway.ListenAndServe(":3000", s.Router))

}
