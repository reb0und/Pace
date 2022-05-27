package cmv2

import (
	"github.com/Nano-Software/Pace/internal/tasks"
)

func Initialize() {
	_ = tasks.TaskMap{
		INITIALIZE: initialize,
		MONITOR:    monitor,
	}
}
