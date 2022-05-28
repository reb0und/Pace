package tasks

import "errors"

var (
	FinishedState TaskState = "finished"
	ErrorState    TaskState = "error"

	ModuleDoesNotExistError    = errors.New("module does not exist")
	TaskStateDoesNotExistError = errors.New("task state does not exist")

	moduleMap = make(map[string]*Module)
)

func RegisterModule(moduleName string) *Module {
	moduleMap[moduleName] = &Module{}

	return moduleMap[moduleName]
}

func DoesModuleExist(module string) bool {
	_, ok := moduleMap[module]

	return ok
}

func (m *Module) DoesStateExist(state TaskState) bool {
	_, ok := m.taskMap[state]

	return ok
}

func (m *Module) AddHandlers(taskMap *TaskMap) {
	m.taskMap = make(TaskMap)
	for state, handler := range *taskMap {
		m.taskMap[state] = handler
	}
}

func (m *Module) SetInitialState(state TaskState) error {
	if m.DoesStateExist(state) {
		m.initialState = state
		return nil
	}

	return TaskStateDoesNotExistError
}
