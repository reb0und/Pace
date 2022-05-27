package engine

import "github.com/Nano-Software/Pace/internal/tasks"

func StartTask() {
	//	call first task method in interface?
}

func StopTask(t *tasks.Task) {
	t.Cancel()
}
