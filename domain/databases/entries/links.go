package entries

type links struct {
	list []Link
}

func createLinks(
	list []Link,
) Links {
	out := links{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *links) List() []Link {
	return obj.list
}
