package spec

import (
	tasks2 "github.com/eternalblue/monsta/internal/tasks"
)

type Spec interface {
	Tasks() (*[]tasks2.Task, error)
}
