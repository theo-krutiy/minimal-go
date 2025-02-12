package server

func (s *Server) routes() {
	s.r.HandleFunc("POST /users", s.handleCreateUser())
	s.r.HandleFunc("POST /authenticate", s.handleAuthenticate())
	s.r.HandleFunc("GET /items", s.handleGetItems())
}
