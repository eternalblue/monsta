package tasks

import "github.com/eternalblue/monsta/pkg/commands"

type Task interface {
	GetCommands() *[]commands.Command
	Name() string
}

type TaskImpl struct {
	Commands *[]commands.Command
	TaskName string
}

func (t TaskImpl) Name() string {
	return t.TaskName
}

func (t TaskImpl) GetCommands() *[]commands.Command {
	return t.Commands
}
