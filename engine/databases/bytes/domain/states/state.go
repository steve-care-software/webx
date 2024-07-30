package states

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"

type state struct {
	isDeleted  bool
	containers containers.Containers
}

func createState() State {
	return createStateInternally(false, nil)
}

func createStateWithContainers(
	containers containers.Containers,
) State {
	return createStateInternally(false, containers)
}

func createStateWithDeleted(
	containers containers.Containers,
) State {
	return createStateInternally(true, containers)
}

func createStateWithContainersAndDeleted(
	containers containers.Containers,
) State {
	return createStateInternally(true, containers)
}

func createStateInternally(
	isDeleted bool,
	containers containers.Containers,
) State {
	out := state{
		isDeleted:  isDeleted,
		containers: containers,
	}

	return &out
}

// IsDeleted returns true if deleted, false otherwise
func (obj *state) IsDeleted() bool {
	return obj.isDeleted
}

// HasContainers returns true if there is containers, false otherwise
func (obj *state) HasContainers() bool {
	return obj.containers != nil
}

// Containers returns the containers, if any
func (obj *state) Containers() containers.Containers {
	return obj.containers
}
