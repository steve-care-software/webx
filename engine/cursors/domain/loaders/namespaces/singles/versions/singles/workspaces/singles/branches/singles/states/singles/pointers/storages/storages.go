package storages

type storages struct {
	list []Storage
}

func createStorages(
	list []Storage,
) Storages {
	out := storages{
		list: list,
	}

	return &out
}

// List returns the list of storages
func (obj *storages) List() []Storage {
	return obj.list
}
