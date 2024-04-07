package main

import (
	"github.com/youth-design/fit_app/libs/logger"
	"github.com/youth-design/fit_app/services/auth/internal/grpc/server"
)

func main() {
	log := logger.Get()

	s := server.New(log, 3030)
	err := s.Start()
	if err != nil {
		panic(err)
	}
}
