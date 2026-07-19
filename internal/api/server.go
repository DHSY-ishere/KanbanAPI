package api

import (
	"net/http"

	"github.com/DHSY-ishere/kanbanAPI/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	cfg    config.Config
	router *chi.Mux
}

func NewServer(cfg config.Config) *Server {
	s := &Server{
		cfg:    cfg,
		router: chi.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", s.handleHealth)

	})
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.cfg.Addr, s.router)
}
