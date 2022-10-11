package trees

type contents struct {
	list []Content
}

func createContents(
	list []Content,
) Contents {
	out := contents{
		list: list,
	}

	return &out
}

// List represents the list of contents
func (obj *contents) List() []Content {
	return obj.list
}
