package contents

type contents struct {
	mp   map[string]Content
	list []Content
}

func createContents(
	mp map[string]Content,
	list []Content,
) Contents {
	out := contents{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the contents
func (obj *contents) List() []Content {
	return obj.list
}
