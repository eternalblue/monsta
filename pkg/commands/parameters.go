package commands

import (
	"errors"
	"fmt"
)

var ErrBadStrAssertion = errors.New("command parameter value is not a string")

type Parameters map[string]interface{}

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
