package spec

import "go.uber.org/zap"

func Run(spec Spec) error {
	tasks, err := spec.Tasks()
	if err != nil {
		return err
	}

	for _, task := range *tasks {
		zap.L().Info("running task", zap.String("task", task.Name()))

		var lastOutput *string

		for _, command := range *task.GetCommands() {
			zap.L().Info("running command", zap.String("type", command.Type()), zap.Any("command", command))
			err := command.ValidateInput(lastOutput)
			if err != nil {
				return err
			}

			lastOutput, err = command.Execute(lastOutput)
			if err != nil {
				return err
			}

			zap.L().Info("finishd command", zap.String("type", command.Type()), zap.Any("output", lastOutput))
		}
	}

	return nil
}
