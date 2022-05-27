package tasks

import "context"

type TaskState string

type TaskMap map[TaskState]interface{}

type Task struct {
	TaskIsActive bool
	Module       string
	Input        string
	Internal     interface{}
	Context      context.Context
	Cancel       context.CancelFunc
	Secret       string
}

type Module struct {
	initialState TaskState
	taskMap      TaskMap
}
