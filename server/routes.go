package server

func (s *Server) routes() {
	s.router.Route("/poi", s.routePOI)
}
