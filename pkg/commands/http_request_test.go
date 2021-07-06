package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpRequestCommand_FromParams(t *testing.T) {
	t.Run("Setup works correctly", func(t *testing.T) {
		params := map[string]interface{}{
			"endpoint": "https://api.google.com",
			"method":   "GET",
		}

		cmd := HttpRequestCommand{}

		err := cmd.Setup(params)

		assert.NoError(t, err, "http request command Setup threw error => %s", err)
		assert.Equal(t, "GET", cmd.Method)
	})
}
