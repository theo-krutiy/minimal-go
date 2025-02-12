package server

import (
	"net/http"

	"github.com/theo-krutiy/minimal-go/internal/auth"
	"github.com/theo-krutiy/minimal-go/internal/shop"
)

type Server struct {
	s  *http.Server
	r  *http.ServeMux
	Db interface {
		auth.Database
		shop.Database
	}
	secret []byte
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
