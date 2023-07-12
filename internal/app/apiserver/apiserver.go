package apiserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *http.ServeMux //*mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: http.NewServeMux(), //mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	err := s.configureLogger()

	if err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.HandleHome())
	s.router.HandleFunc("/caesar/", s.HandleCaesar())
	s.router.HandleFunc("/vigenere/", s.HandleVigenere())
	s.router.HandleFunc("/simplesubstitution/", s.HandleSimpleSC())
	s.router.HandleFunc("/affine/", s.HandleAffine())

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	s.router.Handle("/static/", http.StripPrefix("/static", fileServer))
}
