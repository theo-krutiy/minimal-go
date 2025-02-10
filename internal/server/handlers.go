package server

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) handleGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You GET it, right?")
}

func (s *Server) handlePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You've POSTed this:\n\n")
	io.Copy(w, r.Body)
	fmt.Fprintf(w, "\n\nYour address is %v", r.RemoteAddr)
}
