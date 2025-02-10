package server

func (s *Server) routes() {
	s.r.HandleFunc("POST /users", s.handleCreateUser())
}
