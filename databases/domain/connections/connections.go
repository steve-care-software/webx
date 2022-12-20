package connections

import (
	"errors"
	"fmt"
)

type connections struct {
	mp   map[uint]Connection
	list []Connection
}

func createConnections(
	mp map[uint]Connection,
	list []Connection,
) Connections {
	out := connections{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the connections
func (obj *connections) List() []Connection {
	return obj.list
}

// Fetch fetches a connection by its identifier
func (obj *connections) Fetch(identifier uint) (Connection, error) {
	if ins, ok := obj.mp[identifier]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no Connection for the given identifier: %d", identifier)
	return nil, errors.New(str)
}
