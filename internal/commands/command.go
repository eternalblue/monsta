// Package commands are the base block to build Tasks
package commands

import (
	"github.com/eternalblue/monsta/pkg/environment"
)

type Command interface {
	// Setup configures a command from an environment.Environment
	Setup(env environment.Environment) error
	// Execute a command for a given input. Returns a string as output or an error if something fails.
	Execute(input *string) (*string, error)
	ValidateInput(input *string) error
	// Type returns the type of the command used to parse the spec: http_request, s3cpy, etc...
	Type() string
}
