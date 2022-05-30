package engine

import (
	"github.com/Nano-Software/Pace/internal/tasks"
)

func StartTask(t *tasks.Task) {
	// TODO: create loop for each task state and call each then handle it with an if statement depending on if task is active
	t.TaskIsActive = true
}

func HandleState() {}

func StopTask(t *tasks.Task) {
	t.TaskIsActive = false
	t.Internal = nil
}
