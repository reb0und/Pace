package cmv2

import (
	"github.com/Nano-Software/Pace/internal/tasks"
)

func Initialize() error {
	cmv2Module := tasks.RegisterModule("cmv2")

	cmv2Module.AddHandlers(&tasks.TaskMap{
		INITIALIZE: initialize,
		MONITOR:    monitor,
	})

	if err := cmv2Module.SetInitialState(INITIALIZE); err != nil {
		return err
	}

	return nil
}
