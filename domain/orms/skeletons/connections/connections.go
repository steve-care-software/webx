package connections

import (
	"errors"
	"fmt"
)

type connections struct {
	mp   map[string]Connection
	list []Connection
}

func createConnections(
	mp map[string]Connection,
	list []Connection,
) Connections {
	out := connections{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the list
func (obj *connections) List() []Connection {
	return obj.list
}

// Fetch fetches a connection by name
func (obj *connections) Fetch(name string) (Connection, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no Connection named '%s'", name)
	return nil, errors.New(str)
}
