// Package commands are the base block to build Tasks
package commands

type Command interface {
	// Setup configures a command from a map for parameters.
	//
	//i.e. FromParams(map{"type": "http_request", "endpoint": "https://jondoe.com", "method": "GET"})
	Setup(params Parameters) error
	Execute(input *string) (*string, error)
	ValidateInput(input *string) error
	// Type returns the type of the command used to parse the spec: http_request, s3cpy, etc...
	Type() string
}
