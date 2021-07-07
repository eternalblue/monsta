package commands

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aymerick/raymond"
	"github.com/eternalblue/monsta/pkg/environment"
)

const templateCommandType = "format"

type TemplateCommand struct {
	Template string `validate:"required"`
	Encoding string `validate:"required,oneof=plain base64"`
}

func (cmd *TemplateCommand) Setup(env environment.Environment) error {
	if cmd.Encoding == "base64" {
		decoded, err := base64.StdEncoding.DecodeString(cmd.Template)
		if err != nil {
			return err
		}

		cmd.Template = string(decoded)
	}

	return nil
}

func (cmd TemplateCommand) Execute(input *string) (*string, error) {
	jsonMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(*input), &jsonMap)
	if err != nil {
		return nil, err
	}

	render, err := raymond.Render(cmd.Template, jsonMap)
	if err != nil {
		return nil, err
	}

	return &render, nil
}

func (cmd TemplateCommand) ValidateInput(input *string) error {
	return nil
}

func (cmd TemplateCommand) Type() string {
	return templateCommandType
}
