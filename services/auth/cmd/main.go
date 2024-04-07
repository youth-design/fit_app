package main

import "github.com/youth-design/fit_app/libs/logger"

func main() {
	log := logger.Get()
	log.Debug().Msg("Hello")
}
