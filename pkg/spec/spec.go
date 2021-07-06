package spec

import "github.com/eternalblue/monsta/pkg/tasks"

type Spec interface {
	Tasks() (*[]tasks.Task, error)
}
