package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/Nano-Software/Pace/configs"
	"github.com/Nano-Software/Pace/internal/bootstrap"
	"github.com/Nano-Software/Pace/internal/modules/cmv2"
)

func main() {
	if err := bootstrap.InitializeZapLogger(); err != nil {
		log.Fatalln(err.Error())
	}

	_, err := configs.LoadConfig("./config.yml")
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	cmv2.Initialize()

	// TODO: Create function or some sort of handler here for config
}
