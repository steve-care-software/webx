package criterias

type criteria struct {
	name            string
	index           uint
	includeChannels bool
	content         Content
}

func createCriteria(
	name string,
	index uint,
	includeChannels bool,
) Criteria {
	return createCriteriaInternally(name, index, includeChannels, nil)
}

func createCriteriaWithContent(
	name string,
	index uint,
	includeChannels bool,
	content Content,
) Criteria {
	return createCriteriaInternally(name, index, includeChannels, content)
}

func createCriteriaInternally(
	name string,
	index uint,
	includeChannels bool,
	content Content,
) Criteria {
	out := criteria{
		name:            name,
		index:           index,
		includeChannels: includeChannels,
		content:         content,
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

// IncludeChannels returns true if channels are included, false otherwise
func (obj *criteria) IncludeChannels() bool {
	return obj.includeChannels
}

// HasContent returns true if there is a content, false otherwise
func (obj *criteria) HasContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *criteria) Content() Content {
	return obj.content
}
