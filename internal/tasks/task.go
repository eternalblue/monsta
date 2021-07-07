package tasks

import (
	commands2 "github.com/eternalblue/monsta/internal/commands"
)

type Task interface {
	GetCommands() *[]commands2.Command
	Name() string
}

type TaskImpl struct {
	Commands *[]commands2.Command
	TaskName string
}

func (t TaskImpl) Name() string {
	return t.TaskName
}

func (t TaskImpl) GetCommands() *[]commands2.Command {
	return t.Commands
}
