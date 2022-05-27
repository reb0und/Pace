package tasks

import "github.com/Nano-Software/Pace/configs"

func InitializeTasks(config *configs.Config) error {
	if !DoesModuleExist(config.Module) {
		return ModuleDoesNotExistError
	}

	return nil
}
