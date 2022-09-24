package criterias

type criteria struct {
	name    string
	index   uint
	content Content
}

func createCriteria(
	name string,
	index uint,
	content Content,
) Criteria {
	out := criteria{
		name:    name,
		index:   index,
		content: content,
	}

	return &out
}

// Name returns the name
func (obj *criteria) Name() string {
	return obj.name
}

// Index returns the index
func (obj *criteria) Index() uint {
	return obj.index
}

// Content returns the content
func (obj *criteria) Content() Content {
	return obj.content
}
