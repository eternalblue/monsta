// just a util to return structs from strings :)

package commands

import (
	"fmt"
	"reflect"
)

// commandRegistry maps a type name to the actual type.
type commandRegistry map[string]reflect.Type

var Registry = make(commandRegistry)

func init() {
	Registry.set(new(HTTPRequestCommand))
	Registry.set(new(S3CpyCommand))
	Registry.set(new(TemplateCommand))
}

func (t commandRegistry) set(i Command) {
	t[i.Type()] = reflect.TypeOf(i).Elem()
}

// GetInstance for a given type string the actual type.
func (t commandRegistry) GetInstance(name string) (Command, error) {
	if typ, ok := t[name]; ok {
		if tc, ok := reflect.New(typ).Interface().(Command); ok {
			return tc, nil
		}
		return nil, fmt.Errorf("cannot assert %s as Command", name)
	}
	return nil, fmt.Errorf("command %s does not exists", name)
}
