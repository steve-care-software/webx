package databases

type databases struct {
	list []Database
}

func createDatabases(
	list []Database,
) Databases {
	out := databases{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *databases) List() []Database {
	return obj.list
}
