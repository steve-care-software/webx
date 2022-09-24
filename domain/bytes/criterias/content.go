package criterias

type content struct {
	requirement []uint
	child       Criteria
}

func createContentWithRequirement(
	requirement []uint,
) Content {
	return createContentInternally(requirement, nil)
}

func createContentWithChild(
	child Criteria,
) Content {
	return createContentInternally(nil, child)
}

func createContentInternally(
	requirement []uint,
	child Criteria,
) Content {
	out := content{
		requirement: requirement,
		child:       child,
	}

	return &out
}

// IsRequirement returns true if there is a requirement, false otherwise
func (obj *content) IsRequirement() bool {
	return obj.requirement != nil
}

// Requirement returns the requirement, if any
func (obj *content) Requirement() []uint {
	return obj.requirement
}

// IsChild returns true if there is a child, false otherwise
func (obj *content) IsChild() bool {
	return obj.child != nil
}

// Child returns the child, if any
func (obj *content) Child() Criteria {
	return obj.child
}
