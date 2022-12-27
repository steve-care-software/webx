package trees

type value struct {
	content byte
	prefix  Trees
}

func createValue(
	content byte,
) Value {
	return createValueInternally(content, nil)
}

func createValueWithPrefix(
	content byte,
	prefix Trees,
) Value {
	return createValueInternally(content, prefix)
}

func createValueInternally(
	content byte,
	prefix Trees,
) Value {
	out := value{
		content: content,
		prefix:  prefix,
	}

	return &out
}

// Content returns the content
func (obj *value) Content() byte {
	return obj.content
}

// HasPrefix returns true if there is a prefix, false otherwise
func (obj *value) HasPrefix() bool {
	return obj.prefix != nil
}

// Prefix returns the prefix, if any
func (obj *value) Prefix() Trees {
	return obj.prefix
}
