package main

import (
	"log"

	"github.com/Nano-Software/Pace/modules/cmv2"

	"go.uber.org/zap"

	"github.com/Nano-Software/Pace/configs"
	"github.com/Nano-Software/Pace/internal/bootstrap"
	"github.com/Nano-Software/Pace/internal/tasks"
)

func main() {
	// Initialize Uber Zap Logger
	if err := bootstrap.InitializeZapLogger(); err != nil {
		log.Fatalln(err.Error())
	}

	// Load Task Configuration
	config, err := configs.LoadConfig("./config.yml")
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	// Initialize CMV2 Module
	if err := cmv2.Initialize(); err != nil {
		zap.L().Fatal(err.Error())
	}

	if err := tasks.InitializeTasks(config); err != nil {
		zap.L().Fatal(err.Error())
	}

	// TODO: Create overarching task controller with monitor and other data such as block hash
}
