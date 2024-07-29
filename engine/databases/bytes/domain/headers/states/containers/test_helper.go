package containers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers/pointers"

// NewContainersForTests creates a new containers for tests
func NewContainersForTests(list []Container) Containers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContainerForTests creates a new container for tests
func NewContainerForTests(keyname string, pointers pointers.Pointers) Container {
	ins, err := NewContainerBuilder().Create().WithKeyname(keyname).WithPointers(pointers).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
