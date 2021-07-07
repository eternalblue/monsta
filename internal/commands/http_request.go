package commands

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eternalblue/monsta/pkg/environment"
	"github.com/eternalblue/monsta/pkg/utils"
)

const commandName = "http_request"

var ErrNilInput = errors.New("input cannot be nil for http_request with POST Method")

// HTTPRequestCommand implementation.
type HTTPRequestCommand struct {
	client   *http.Client
	Endpoint string `json:"endpoint"`
	Method   string `validate:"required,oneof=GET POST PUT DELETE PATCH" json:"method"`
	Body     string
}

// Setup ...
func (cmd *HTTPRequestCommand) Setup(env environment.Environment) error {
	cmd.client = env.NetClient()

	return nil
}

// Execute an HTTPRequestCommand.
func (cmd HTTPRequestCommand) Execute(input *string) (*string, error) {
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

func (HTTPRequestCommand) ValidateInput(input *string) error {
	return nil
}

// Type returns the type of HTTPRequestCommand.
func (cmd HTTPRequestCommand) Type() string {
	return commandName
}
