package criterias

type criteria struct {
	name            string
	index           uint
	includeChannels bool
	child           Criteria
}

func createCriteria(
	name string,
	index uint,
	includeChannels bool,
) Criteria {
	return createCriteriaInternally(name, index, includeChannels, nil)
}

func createCriteriaWithChild(
	name string,
	index uint,
	includeChannels bool,
	child Criteria,
) Criteria {
	return createCriteriaInternally(name, index, includeChannels, child)
}

func createCriteriaInternally(
	name string,
	index uint,
	includeChannels bool,
	child Criteria,
) Criteria {
	out := criteria{
		name:            name,
		index:           index,
		includeChannels: includeChannels,
		child:           child,
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

// HasChild returns true if there is a child, false otherwise
func (obj *criteria) HasChild() bool {
	return obj.child != nil
}

// Child returns the child, if any
func (obj *criteria) Child() Criteria {
	return obj.child
}
