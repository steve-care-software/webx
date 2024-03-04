package connections

import (
	"errors"
	"fmt"
	"strings"
)

type connections struct {
	mpByPaths map[string]Connection
	mp        map[string]Connection
	list      []Connection
}

func createConnections(
	mpByPaths map[string]Connection,
	mp map[string]Connection,
	list []Connection,
) Connections {
	out := connections{
		mpByPaths: mpByPaths,
		mp:        mp,
		list:      list,
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

// FetchByPaths fetches by paths
func (obj *connections) FetchByPaths(from []string, to []string) (Connection, error) {
	keyname := createKeynameFromPaths(from, to)
	if ins, ok := obj.mpByPaths[keyname]; ok {
		return ins, nil
	}

	fromStr := strings.Join(from, "/")
	toStr := strings.Join(to, "/")
	str := fmt.Sprintf("there is no Connection related to the provided paths (from: %s, to: %s)", fromStr, toStr)
	return nil, errors.New(str)
}
