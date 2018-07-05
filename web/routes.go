package web

//go:generate go-bindata -pkg web -o template.go template/

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/bkono/sls-gojo"
	"github.com/guregu/dynamo"
	"github.com/segmentio/ksuid"
)

var (
	indexTmpl  = template.New("index.html")
	defaultTTL = 24 * time.Hour
)

type indexData struct {
	BaseURL   string
	CanCreate bool
	Msg       string
}

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
	s.renderForm(w, r)
}

func (s *Server) create() http.HandlerFunc {
	var (
		init  sync.Once
		sess  *session.Session
		db    *dynamo.DB
		table dynamo.Table
		kmscl *kms.KMS
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			sess = session.Must(session.NewSession())
			db = dynamo.New(sess)
			table = db.Table(s.tableName)
			kmscl = kms.New(sess)
		})

		msg := r.FormValue("msg")
		if len(msg) == 0 {
			s.renderMsg(w, r, "msg is required")
		}

		msgBytes, err := encrypt(kmscl, s.kmsKeyID, []byte(msg))
		if err != nil {
			log.Printf("failed encrypting entry, err = %v\n", err)
			s.renderMsg(w, r, "failed saving entry")
			return
		}

		e := &justonce.Entry{
			ID:        ksuid.New().String(),
			Msg:       msgBytes,
			CreatedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(defaultTTL).Unix(),
		}

		err = table.Put(e).Run()
		if err != nil {

			log.Printf("failed saving entry, err = %v\n", err)
			s.renderMsg(w, r, "failed saving entry")
			return
		}

		base := s.baseURL(r)
		s.renderMsg(w, r, base+"/get/"+e.ID)
	}
}

func (s *Server) get() http.HandlerFunc {
	var (
		init  sync.Once
		sess  *session.Session
		db    *dynamo.DB
		table dynamo.Table
		kmscl *kms.KMS
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			sess = session.Must(session.NewSession())
			db = dynamo.New(sess)
			table = db.Table(s.tableName)
			kmscl = kms.New(sess)
		})

		id := Param(r.Context(), "id")
		var entry justonce.Entry
		err := table.Get("id", id).One(&entry)
		if err == dynamo.ErrNotFound {
			s.renderMsg(w, r, "msg already removed")
			return
		}
		if err != nil {
			log.Printf("failed retrieving entry, err = %v\n", err)
			s.renderMsg(w, r, "msg not retrievable")
			return
		}

		if time.Now().Unix() > entry.ExpiresAt {
			s.renderMsg(w, r, "msg has expired")
			deleteID(table, id)
			return
		}

		txt, err := decrypt(kmscl, entry.Msg)
		if err != nil {
			log.Printf("failed decrypting entry, err = %v\n", err)
			s.renderMsg(w, r, "msg not retrievable")
			return
		}

		s.renderMsg(w, r, string(txt))
		deleteID(table, id)
	}
}

func deleteID(table dynamo.Table, id string) {
	err := table.Delete("id", id).Run()
	if err != nil {
		log.Println("err deleting entry", err)
		// queue up another attempt?
	}
}

func (s *Server) renderForm(w http.ResponseWriter, r *http.Request) {
	s.renderIndex(w, r, "", true)
}

func (s *Server) renderMsg(w http.ResponseWriter, r *http.Request, msg string) {
	s.renderIndex(w, r, msg, false)
}

func (s *Server) renderIndex(w http.ResponseWriter, r *http.Request, msg string, canCreate bool) {
	data := &indexData{
		BaseURL:   s.baseURL(r),
		Msg:       msg,
		CanCreate: canCreate,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	indexTmpl.Execute(w, data)
}

func (s *Server) requireAWSEnvs(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(s.tableName) == 0 || len(s.kmsKeyID) == 0 {
			http.Error(w, "table and kms key required", http.StatusInternalServerError)
			return
		}

		h(w, r)
	}
}

var (
	bots = []string{
		"Slack", "Facebot", "Twitter", "facebook", "Applebot", "LinkedInBot", "SkypeUriPreview",
		"Googlebot", "AdsBot", "Feedfetcher", "bingbot", "Slurp", "msnbot",
	}
)

func (s *Server) ignoreTheBots(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received, ua", r.UserAgent())
		for _, bot := range bots {
			if strings.Contains(r.UserAgent(), bot) {
				log.Println("ua includes a bot, short circuiting request", bot)
				w.Write([]byte(bot + " preview not available"))
				return
			}
		}

		h(w, r)
	}
}

func (s *Server) routes() {
	s.Router.HandleFunc("GET", "/", s.ignoreTheBots(s.index))
	s.Router.HandleFunc("GET", "/home", s.ignoreTheBots(s.index))
	s.Router.HandleFunc("POST", "/create", s.ignoreTheBots(s.requireAWSEnvs(s.create())))
	s.Router.HandleFunc("GET", "/get/:id", s.ignoreTheBots(s.requireAWSEnvs(s.get())))
	s.Router.HandleFunc("GET", "/public/:asset", s.public)
}
