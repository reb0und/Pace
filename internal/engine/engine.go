package engine

import (
	"fmt"

	"github.com/Nano-Software/Pace/internal/tasks"
)

func StartTask(t *tasks.Task) {
	fmt.Println(t)
	//	call first task method in interface?
}

func StopTask(t *tasks.Task) {
	t.Cancel()
}
