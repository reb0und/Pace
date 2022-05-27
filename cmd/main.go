package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/Nano-Software/Pace/configs"
	"github.com/Nano-Software/Pace/internal/bootstrap"
	"github.com/Nano-Software/Pace/internal/modules/cmv2"
)

func main() {
	// Initialize Uber Zap Logger
	if err := bootstrap.InitializeZapLogger(); err != nil {
		log.Fatalln(err.Error())
	}

	// Load Task Configuration
	if _, err := configs.LoadConfig("./config.yml"); err != nil {
		zap.L().Fatal(err.Error())
	}

	// Initialize CMV2 Module
	if err := cmv2.Initialize(); err != nil {
		zap.L().Fatal(err.Error())
	}

	// TODO: Create function or some sort of handler here for config
}
