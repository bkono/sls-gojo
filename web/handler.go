package web

//go:generate go-bindata -pkg web -o template.go template/

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bkono/sls-gojo"
	"github.com/guregu/dynamo"
	"github.com/segmentio/ksuid"
)

var (
	indexTmpl = template.New("index.html")
)

func init() {
	index, err := Asset("template/index.html")
	if err != nil {
		panic(err)
	}

	template.Must(indexTmpl.Parse(string(index)))
}

func public(w http.ResponseWriter, r *http.Request) {
	name := Param(r.Context(), "asset")
	asset, err := Asset("template/" + name)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch {
	case strings.HasSuffix(name, ".css"):
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case strings.HasSuffix(name, ".html"):
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	default:
		w.Header().Set("Content-Type", "text/plain")
	}

	w.Write(asset)
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	stage := os.Getenv("STAGE")
	data := struct {
		Stage    string
		HasStage bool
	}{
		Stage:    stage,
		HasStage: len(stage) > 0,
	}
	indexTmpl.Execute(w, data)
}

func create(w http.ResponseWriter, r *http.Request) {
	tableName := os.Getenv("TABLE_NAME")
	if len(tableName) == 0 {
		http.Error(w, "table not configured", http.StatusInternalServerError)
		return
	}

	msg := r.FormValue("msg")
	if len(msg) == 0 {
		http.Error(w, "msg is required", 400)
	}

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess)
	table := db.Table(tableName)
	e := &justonce.Entry{
		ID:        ksuid.New().String(),
		Msg:       msg,
		TTL:       120,
		CreatedAt: time.Now(),
	}

	err := table.Put(e).Run()
	if err != nil {

		log.Printf("failed saving entry, err = %v\n", err)
		http.Error(w, "failed saving entry", 500)
		return
	}

	base := fmt.Sprintf("https://%s", r.Host)
	stage := os.Getenv("STAGE")
	if len(stage) > 0 {
		base = fmt.Sprintf("%s/%s", base, stage)
	}

	log.Println(base)

	w.Write([]byte(fmt.Sprintf("%s/get/%s", base, e.ID)))
}

func get(w http.ResponseWriter, r *http.Request) {
	// extract this, context maybe? middleware?
	tableName := os.Getenv("TABLE_NAME")
	if len(tableName) == 0 {
		http.Error(w, "table not configured", http.StatusInternalServerError)
		return
	}
	id := Param(r.Context(), "id")

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess)
	table := db.Table(tableName)
	var entry justonce.Entry
	err := table.Get("id", id).One(&entry)
	if err != nil {
		log.Printf("failed retrieving entry, err = %v\n", err)
		http.Error(w, "failed retrieving entry", 500)
		return
	}

	w.Write([]byte(entry.Msg))
}

func NewHandler() *Router {
	router := NewRouter()
	router.HandleFunc("GET", "/home", index)
	router.HandleFunc("POST", "/create", create)
	router.HandleFunc("GET", "/get/:id", get)
	router.HandleFunc("GET", "/public/:asset", public)
	return router
}
