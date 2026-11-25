package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "", log.LstdFlags)

	s := server.New(logger)

	if err := s.Start(); err != nil {
		logger.Fatal(err)
	}
}
