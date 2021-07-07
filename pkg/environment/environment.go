package environment

import (
	"github.com/eternalblue/monsta/pkg/utils"
	"net/http"
)

// DefaultEnvironment with default settings. For a custom one, implement Environment interface.
var DefaultEnvironment = defaultEnvironment{}

type Environment interface {
	NetClient() *http.Client
	S3Client() *utils.S3Client
}

type defaultEnvironment struct{}

func (d defaultEnvironment) NetClient() *http.Client {
	return http.DefaultClient
}

func (d defaultEnvironment) S3Client() *utils.S3Client {
	return utils.DefaultS3Client
}
