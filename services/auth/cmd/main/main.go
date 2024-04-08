package main

import (
	"fmt"
	"github.com/youth-design/fit_app/libs/logger"
	"github.com/youth-design/fit_app/services/auth/internal/config"
	"github.com/youth-design/fit_app/services/auth/internal/grpc/server"
	"os"
	"os/signal"
	"syscall"
)

const envProd = "prod"

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("config path is not specified"))
	}
	cfg := config.MustRead(os.Args[1])

	logger.Initialize(cfg.Env == envProd)
	log := logger.Get()

	s := server.New(log, 3030)
	err := s.Start()
	if err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	s.GracefulShutdown()

}
