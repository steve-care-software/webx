package keys

type keys struct {
	list []Key
}

func createKeys(
	list []Key,
) Keys {
	out := keys{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *keys) List() []Key {
	return obj.list
}
