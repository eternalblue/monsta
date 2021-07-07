package spec

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/eternalblue/monsta/internal/commands"
	"github.com/eternalblue/monsta/internal/tasks"
	"github.com/eternalblue/monsta/pkg/environment"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

const (
	stepsKey       = "steps"
	commandTypeKey = "type"
)

var (
	validate *validator.Validate

	ErrMissingTypeKey  = errors.New("missing 'type' on command")
	ErrMissingStepsKey = errors.New("missing 'steps' on command")
)

func init() {
	validate = validator.New()
}

// JSON spec format.
type JSON struct {
	// Content is the content of the json in bytes.
	Content     []byte
	Environment environment.Environment
}

// FromJSONFile given a path returns a JSON struct instance.
func FromJSONFile(path string, environment environment.Environment) (*JSON, error) {
	zap.L().Info("loading json", zap.String("path", path))
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	zap.L().Info("working directory", zap.String("wd", wd))

	b, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, path))
	if err != nil {
		return nil, err
	}

	jsonSpec := JSON{Content: b, Environment: environment}

	zap.L().Debug("json loaded", zap.String("content", string(b)))

	return &jsonSpec, nil
}

// FromJSONBytes returns a JSON from a byte array.
func FromJSONBytes(json []byte) *JSON {
	return &JSON{Content: json}
}

// Tasks return an array of tasks by parsing JSON.Content or an error if it fails
func (spec JSON) Tasks() (*[]tasks.Task, error) {
	zap.L().Info("parsing spec tasks")

	var t []tasks.Task

	jsonParsed, err := gabs.ParseJSON(spec.Content)
	if err != nil {
		return nil, err
	}

	for specEntryName, specEntry := range jsonParsed.S().ChildrenMap() {
		zap.L().Debug("iterating json", zap.String("entry name", specEntryName))

		if !specEntry.ExistsP(stepsKey) {
			return nil, ErrMissingStepsKey
		}

		task := tasks.TaskImpl{
			TaskName: specEntryName,
		}

		steps, err := spec.parseSteps(specEntry.S(stepsKey))
		if err != nil {
			return nil, err
		}

		task.Commands = steps

		t = append(t, task)
	}
	return &t, nil
}

func (spec JSON) parseSteps(steps *gabs.Container) (*[]commands.Command, error) {
	var cmds []commands.Command

	for _, step := range steps.Children() {
		zap.L().Info("parsing step", zap.String("step", step.String()))

		if !step.ExistsP(commandTypeKey) {
			return nil, ErrMissingTypeKey
		}

		cmdTypeString := step.Path(commandTypeKey).Data().(string)

		cmdInstance, err := commands.Registry.GetInstance(cmdTypeString)
		if err != nil {
			return nil, err
		}

		var params map[string]interface{}

		err = json.Unmarshal(step.Bytes(), &params)
		if err != nil {
			return nil, err
		}

		err = mapstructure.Decode(params, &cmdInstance)
		if err != nil {
			return nil, err
		}

		err = cmdInstance.Setup(spec.Environment)
		if err != nil {
			return nil, err
		}

		err = validate.Struct(cmdInstance)
		if err != nil {
			return nil, err
		}

		cmds = append(cmds, cmdInstance)
	}

	zap.L().Debug("finished parsing steps", zap.Any("steps", cmds))

	return &cmds, nil
}
