package web

import "os"

type Server struct {
	Router    *Router
	tableName string
	stage     string
}

func (s *Server) hasStage() bool {
	return len(s.stage) > 0
}

func NewServer() *Server {
	s := &Server{
		Router:    NewRouter(),
		tableName: os.Getenv("TABLE_NAME"),
		stage:     os.Getenv("STAGE"),
	}
	s.routes()
	return s
}
