package server

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"http-rest-api/app/config"
	"io"
	"net/http"
)

type APIServer struct {
	config *config.Config
	logger *log.Logger
	router *mux.Router
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		logger: log.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server")
	s.ConfigureRouter()
	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIServer) ConfigureLogger() error {
	level, err := log.ParseLevel(s.config.LogLevel)
	if err != nil {
		s.logger.Error("Could nor parse log level", s.config.LogLevel)
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) ConfigureRouter() {
	s.router.Handle("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!")
	}
}
