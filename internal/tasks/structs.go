package tasks

type TaskState string

type TaskMap map[TaskState]interface{}

type Task struct {
	TaskIsActive bool
	Module       string
	Input        string
	Internal     interface{}
	Secret       string
	Monitor      chan interface{}
}

type Module struct {
	initialState TaskState
	taskMap      TaskMap
}
