// Package commands are the base block to build Tasks
package commands

import (
	"errors"
	"fmt"
)

var ErrBadStrAssertion = errors.New("command parameter value is not a string")

type (
	Parameters map[string]interface{}

	Command interface {
		// Setup configures a command from a map for parameters.
		//
		//i.e. FromParams(map{"type": "http_request", "endpoint": "https://jondoe.com", "method": "GET"})
		Setup(params Parameters) error
		Execute(input *string) (*string, error)
		ValidateInput(input *string) error
		// Type returns the type of the command used to parse the spec: http_request, s3cpy, etc...
		Type() string
	}
)

func (p Parameters) GetString(key string, optional ...bool) (string, error) {
	i, exists := p[key]

	if !exists {
		if len(optional) > 0 && optional[0] {
			return "", nil
		}
		return "", fmt.Errorf("key '%s' not present", key)
	}

	v, ok := i.(string)
	if !ok {
		return "", ErrBadStrAssertion
	}

	return v, nil
}
