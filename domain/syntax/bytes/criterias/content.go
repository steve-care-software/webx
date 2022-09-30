package criterias

type content struct {
	child Criteria
	match []byte
}

func createContentWithChild(
	child Criteria,
) Content {
	return createContentInternally(child, nil)
}

func createContentWithMatch(
	match []byte,
) Content {
	return createContentInternally(nil, match)
}

func createContentInternally(
	child Criteria,
	match []byte,
) Content {
	out := content{
		child: child,
		match: match,
	}

	return &out
}

// IsChild returns true if there is a child, false otherwise
func (obj *content) IsChild() bool {
	return obj.child != nil
}

// Child returns the child, if any
func (obj *content) Child() Criteria {
	return obj.child
}

// IsMatch returns true if there is a match, false otherwise
func (obj *content) IsMatch() bool {
	return obj.match != nil
}

// Match returns the match, if any
func (obj *content) Match() []byte {
	return obj.match
}
