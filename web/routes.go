package web

//go:generate go-bindata -pkg web -o template.go template/

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bkono/sls-gojo"
	"github.com/guregu/dynamo"
	"github.com/segmentio/ksuid"
)

var (
	indexTmpl  = template.New("index.html")
	DefaultTTL = 1 * time.Minute
)

// migrate this to within index with sync.Once pattern
func init() {
	index, err := Asset("template/index.html")
	if err != nil {
		panic(err)
	}

	template.Must(indexTmpl.Parse(string(index)))
}

func (s *Server) public(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct {
		Stage    string
		HasStage bool
	}{
		Stage:    s.stage,
		HasStage: s.hasStage(),
	}

	indexTmpl.Execute(w, data)
}

func (s *Server) create() http.HandlerFunc {
	var (
		init  sync.Once
		sess  *session.Session
		db    *dynamo.DB
		table dynamo.Table
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			sess = session.Must(session.NewSession())
			db = dynamo.New(sess)
			table = db.Table(s.tableName)
		})

		msg := r.FormValue("msg")
		if len(msg) == 0 {
			http.Error(w, "msg is required", 400)
		}

		e := &justonce.Entry{
			ID:        ksuid.New().String(),
			Msg:       msg,
			CreatedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(DefaultTTL).Unix(),
		}

		err := table.Put(e).Run()
		if err != nil {

			log.Printf("failed saving entry, err = %v\n", err)
			http.Error(w, "failed saving entry", 500)
			return
		}

		base := fmt.Sprintf("https://%s", r.Host)
		if s.hasStage() {
			base = fmt.Sprintf("%s/%s", base, s.stage)
		}

		w.Write([]byte(fmt.Sprintf("%s/get/%s", base, e.ID)))
	}
}

func (s *Server) get() http.HandlerFunc {
	var (
		init  sync.Once
		sess  *session.Session
		db    *dynamo.DB
		table dynamo.Table
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			sess = session.Must(session.NewSession())
			db = dynamo.New(sess)
			table = db.Table(s.tableName)
		})

		id := Param(r.Context(), "id")
		var entry justonce.Entry
		err := table.Get("id", id).One(&entry)
		if err == dynamo.ErrNotFound {
			w.Write([]byte("msg already removed"))
			return
		}
		if err != nil {
			log.Printf("failed retrieving entry, err = %v\n", err)
			http.Error(w, "failed retrieving entry", 500)
			return
		}

		if time.Now().Unix() > entry.ExpiresAt {
			w.Write([]byte("msg has expired"))

		} else {
			w.Write([]byte(entry.Msg))
		}

		err = table.Delete("id", id).Run()
		if err != nil {
			log.Println("err deleting entry", err)
			// queue up another attempt?
		}
	}
}

func (s *Server) requireTableName(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(s.tableName) == 0 {
			http.Error(w, "table not configured", http.StatusInternalServerError)
			return
		}
		h(w, r)
	}
}

func (s *Server) routes() {
	s.Router.HandleFunc("GET", "/home", s.index)
	s.Router.HandleFunc("POST", "/create", s.requireTableName(s.create()))
	s.Router.HandleFunc("GET", "/get/:id", s.requireTableName(s.get()))
	s.Router.HandleFunc("GET", "/public/:asset", s.public)
}
