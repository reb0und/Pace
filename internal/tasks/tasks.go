package tasks

import (
	"github.com/Nano-Software/Pace/configs"
)

type TaskState string

type TaskMap map[TaskState]interface{}

type Task struct {
	TaskIsActive bool
	Module       string
	Input        string
	Internal     interface{}
	Secret       string
	Monitor      chan string
}

func InitializeTasks(config *configs.Config) error {
	if !DoesModuleExist(config.Module) {
		return ModuleDoesNotExistError
	}

	for i := 0; i < int(config.Count); i++ {
		go func() {
			// TODO: Possibly create instance of task but that may result in unnecessary allocation of memory however we could reassign it as nil
		}()
	}

	return nil
}
