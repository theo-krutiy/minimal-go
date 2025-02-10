package server

import "net/http"

type Server struct {
	s *http.Server
	r *http.ServeMux
}

func New() *Server {
	r := http.NewServeMux()
	s := &Server{s: &http.Server{Handler: r}}
	s.r = r
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
