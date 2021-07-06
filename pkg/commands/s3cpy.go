package commands

import (
	"github.com/eternalblue/monsta/pkg/utils"
	"os"
)

const s3CpyCommandType = "s3cpy"

type S3CpyCommand struct {
	s3Client *utils.S3Client
	Bucket   string `validate:"required"`
	Key      string `validate:"required"`
	Path     string
}

func NewS3CpyCommand(s3Client *utils.S3Client, bucket, key, path string) *S3CpyCommand {
	return &S3CpyCommand{
		s3Client: s3Client,
		Bucket:   bucket,
		Key:      key,
		Path:     path,
	}
}

func (cmd *S3CpyCommand) Setup(params Parameters) error {
	bucket, err := params.GetString("bucket")
	if err != nil {
		return err
	}

	key, err := params.GetString("key")
	if err != nil {
		return err
	}

	cmd.s3Client = utils.DefaultS3Client
	cmd.Bucket = bucket
	cmd.Key = key

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
	} else {
		content, err := cmd.s3Client.GetContent(cmd.Bucket, cmd.Key)

		if err != nil {
			return nil, err
		}

		return &content, nil
	}
}

func (cmd S3CpyCommand) ValidateInput(input *string) error {
	return nil
}

func (cmd S3CpyCommand) Type() string {
	return s3CpyCommandType
}
