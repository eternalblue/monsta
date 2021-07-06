package commands

const templateCommandType = "template"

type TemplateCommand struct {
	Template string
}

func (t TemplateCommand) Setup(params Parameters) error {
	panic("implement me")
}

func (t TemplateCommand) Execute(input *string) (*string, error) {
	panic("implement me")
}

func (t TemplateCommand) ValidateInput(input *string) error {
	panic("implement me")
}

func (t TemplateCommand) Type() string {
	return templateCommandType
}
