package main

import (
	"github.com/eternalblue/monsta/pkg/spec"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)

	s, err := spec.FromJSONFile("foo.json")
	if err != nil {
		panic(err)
	}

	err = spec.Run(s)
	if err != nil {
		panic(err)
	}
}
