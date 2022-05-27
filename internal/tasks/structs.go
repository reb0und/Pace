package tasks

import "context"

type TaskState string

type TaskMap map[TaskState]interface{}

type Task struct {
	TaskIsActive bool
	Input        string
	Internal     struct{}
	Context      context.Context
	Cancel       context.CancelFunc
	Secret       string
}

type Module struct {
	initialState TaskState
	taskMap      TaskMap
}
