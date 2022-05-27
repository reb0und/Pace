package main

import (
	"log"

	"github.com/Nano-Software/Pace/internal/engine"
	"github.com/Nano-Software/Pace/internal/tasks"

	"go.uber.org/zap"

	"github.com/Nano-Software/Pace/configs"
	"github.com/Nano-Software/Pace/internal/bootstrap"
)

func main() {
	if err := bootstrap.InitializeZapLogger(); err != nil {
		log.Fatalln(err.Error())
	}

	config, err := configs.LoadConfig("./config.yml")
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	switch config.Module {
	case "cmv2":
		for i := 0; i < int(config.Count); i++ {
			engine.StartTask(&tasks.Task{
				Module:   &tasks.Module{},
				Input:    "",
				Internal: struct{}{},
				Context:  nil,
				Cancel:   nil,
				Secret:   "",
			})
		}

	case "":
		zap.L().Fatal("No Module Found")
	default:
		zap.L().Fatal("No Module Found1")

	}
}
