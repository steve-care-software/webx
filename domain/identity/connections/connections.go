package connections

import (
	"bytes"

	uuid "github.com/satori/go.uuid"
)

type connections struct {
	list []Connection
}

func createConnections(
	list []Connection,
) Connections {
	out := connections{
		list: list,
	}

	return &out
}

// List returns the connections
func (obj *connections) List() []Connection {
	return obj.list
}

// ListExcept returns the connections except the one from the given id
func (obj *connections) ListExcept(id uuid.UUID) []Connection {
	output := []Connection{}
	for _, oneConnection := range obj.list {
		if bytes.Compare(oneConnection.ID().Bytes(), id.Bytes()) == 0 {
			continue
		}

		output = append(output, oneConnection)
	}

	return output
}
