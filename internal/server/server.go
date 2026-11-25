package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server представляет HTTP-сервер.
type Server struct {
	logger *log.Logger
	server *http.Server
}

// New создает новый сервер, настраивает маршрутизатор и регистрирует хендлеры.
func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: s,
	}
}

// Start запускает сервер.
func (s *Server) Start() error {
	s.logger.Println("Запускаем сервер")
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	s.logger.Println("Завершаем работу")
	return nil
}
