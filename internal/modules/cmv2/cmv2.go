package cmv2

import (
	"fmt"

	"github.com/Nano-Software/Pace/internal/tasks"
)

func Initialize() {
	taskMap := tasks.TaskMap{
		INITIALIZE: initialize,
		MONITOR:    monitor,
	}
	fmt.Println(taskMap)
}
