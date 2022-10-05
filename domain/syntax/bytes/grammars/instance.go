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

// Name returns the name
func (obj *instance) Name() string {
	if obj.content.IsToken() {
		return obj.content.Token().Name()
	}

	return obj.content.Everything().Name()
}

// Content returns the content
func (obj *instance) Content() InstanceContent {
	return obj.content
}

// IsReverse returns true if reverse, false otherwise
func (obj *instance) IsReverse() bool {
	return obj.isReverse
}
