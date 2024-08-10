package storages

type storagesIns struct {
	list []Storage
}

func createStorages(
	list []Storage,
) Storages {
	out := storagesIns{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *storagesIns) List() []Storage {
	return obj.list
}
