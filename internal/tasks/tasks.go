package tasks

import "github.com/Nano-Software/Pace/configs"

func InitializeTasks(config *configs.Config) error {
	if !DoesModuleExist(config.Module) {
		return ModuleDoesNotExistError
	}

	for i := 0; i < int(config.Count); i++ {
		go func() {}()
	}

	return nil
}
