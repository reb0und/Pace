package tasks

import (
	"github.com/Nano-Software/Pace/configs"
	"github.com/Nano-Software/Pace/internal/engine"
)

func InitializeTasks(config *configs.Config) error {
	if !DoesModuleExist(config.Module) {
		return ModuleDoesNotExistError
	}

	for i := 0; i < int(config.Count); i++ {
		go func() {
			// TODO: Setup context and cancel function here possibly create instance of task but that may result in unnecessary allocation of memory however we could reassign it as nil

			engine.StartTask(&Task{
				TaskIsActive: false,
				Module:       config.Module,
				Input:        config.Input,
				Secret:       config.Secret,
				Monitor:      make(chan interface{}),
			})
		}()
	}

	return nil
}
