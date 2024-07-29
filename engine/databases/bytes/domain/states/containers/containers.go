package containers

import (
	"errors"
	"fmt"
)

type containers struct {
	mp   map[string]Container
	list []Container
}

func createContainers(
	mp map[string]Container,
	list []Container,
) Containers {
	out := containers{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the list of containers
func (obj *containers) List() []Container {
	return obj.list
}

// Fetch fetches a container by keyname
func (obj *containers) Fetch(keyname string) (Container, error) {
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no container for keyname: %s", keyname)
	return nil, errors.New(str)
}
