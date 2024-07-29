package containers

type containers struct {
	list []Container
}

func createContainers(
	list []Container,
) Containers {
	out := containers{
		list: list,
	}

	return &out
}

// List returns the list of containers
func (obj *containers) List() []Container {
	return obj.list
}
