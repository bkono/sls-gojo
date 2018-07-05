package web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	Router    *Router
	tableName string
	stage     string
	kmsKeyID  string
}

func (s *Server) hasStage() bool {
	return len(s.stage) > 0
}

func (s *Server) baseURL(r *http.Request) string {
	base := fmt.Sprintf("https://%s", r.Host)
	// assume that if we aren't on the amazonaws domain, that a base path mapping is not set
	if s.hasStage() && strings.Contains(r.Host, "amazonaws.com") {
		base = base + "/" + s.stage
	}

	return base
}

func NewServer() *Server {
	s := &Server{
		Router:    NewRouter(),
		tableName: os.Getenv("TABLE_NAME"),
		stage:     os.Getenv("STAGE"),
		kmsKeyID:  os.Getenv("KMS_KEY_ID"),
	}

	s.routes()
	return s
}
