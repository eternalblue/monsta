package main

import (
	"github.com/eternalblue/monsta/pkg/environment"
	"github.com/eternalblue/monsta/pkg/spec"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)

	s, err := spec.FromJSONFile("cmd/example_format.json", environment.DefaultEnvironment)
	if err != nil {
		panic(err)
	}

	err = spec.Run(s)
	if err != nil {
		panic(err)
	}
}
