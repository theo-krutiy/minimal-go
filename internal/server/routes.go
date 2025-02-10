package server

func (s *Server) routes() {
	s.r.HandleFunc("GET /example", s.handleGET)
	s.r.HandleFunc("POST /example", s.handlePOST)
}
