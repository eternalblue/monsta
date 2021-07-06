package commands

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eternalblue/monsta/pkg/utils"
)

const commandName = "http_request"

var (
	ErrNilInput = errors.New("input cannot be nil for http_request with POST Method")
)

// HttpRequestCommand implementation.
type HttpRequestCommand struct {
	client   *http.Client
	Endpoint string `validate:"required,url"`
	Method   string `validate:"required,oneof=GET POST PUT DELETE PATCH"`
	Body     string
}

// Setup ...
func (cmd *HttpRequestCommand) Setup(params Parameters) error {
	endpoint, err := params.GetString("endpoint")
	if err != nil {
		return err
	}

	method, err := params.GetString("method")
	if err != nil {
		return err
	}

	body, err := params.GetString("body", true)
	if err != nil {
		return err
	}

	cmd.Endpoint = endpoint
	cmd.Method = method
	cmd.Body = body
	cmd.client = http.DefaultClient

	return nil
}

// Execute an HttpRequestCommand.
func (cmd HttpRequestCommand) Execute(input *string) (*string, error) {
	if cmd.Method != http.MethodGet && cmd.Method != http.MethodDelete {
		if input == nil {
			if cmd.Body == "" {
				return nil, ErrNilInput
			} else {
				input = &cmd.Body
			}
		}
	} else {
		input = utils.StrPointer("")
	}

	request, err := http.NewRequest(cmd.Method, cmd.Endpoint, strings.NewReader(*input))

	if err != nil {
		return nil, err
	}

	response, err := cmd.client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	output := string(body)

	return &output, nil
}

func (HttpRequestCommand) ValidateInput(input *string) error {
	return nil
}

// Type returns the type of HttpRequestCommand.
func (cmd HttpRequestCommand) Type() string {
	return commandName
}
