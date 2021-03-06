package commands

import (
	"os"

	"github.com/eternalblue/monsta/pkg/environment"
	"github.com/eternalblue/monsta/pkg/utils"
)

const s3CpyCommandType = "s3cpy"

type S3CpyCommand struct {
	s3Client *utils.S3Client
	Bucket   string `validate:"required"`
	Key      string `validate:"required"`
	Path     string
}

func (cmd *S3CpyCommand) Setup(environment environment.Environment) error {
	cmd.s3Client = environment.S3Client()

	return nil
}

func (cmd S3CpyCommand) Execute(input *string) (*string, error) {
	if cmd.Path != "" {
		outFile, err := os.Create(cmd.Path)
		if err != nil {
			return nil, err
		}

		err = cmd.s3Client.DownloadFile(cmd.Bucket, cmd.Key, outFile)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	content, err := cmd.s3Client.GetContent(cmd.Bucket, cmd.Key)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (cmd S3CpyCommand) ValidateInput(input *string) error {
	return nil
}

func (cmd S3CpyCommand) Type() string {
	return s3CpyCommandType
}
