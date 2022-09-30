package grammars

type instance struct {
	content   InstanceContent
	isReverse bool
}

func createInstance(
	content InstanceContent,
	isReverse bool,
) Instance {
	out := instance{
		content:   content,
		isReverse: isReverse,
	}

	return &out
}

// Content returns the content
func (obj *instance) Content() InstanceContent {
	return obj.content
}

// IsReverse returns true if reverse, false otherwise
func (obj *instance) IsReverse() bool {
	return obj.isReverse
}
