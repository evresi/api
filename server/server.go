package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	db     *pgxpool.Pool
	router *chi.Mux
}

func NewServer(postgresUrl string) (Server, error) {
	pool, err := pgxpool.Connect(context.Background(), postgresUrl)
	if err != nil {
		return Server{}, err
	}

	router := chi.NewRouter()

	server := Server{pool, router}

	server.routes()

	return server, nil
}

func (s *Server) Serve() {
	http.ListenAndServe(":3030", s.router)
}
