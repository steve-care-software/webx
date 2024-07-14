package begins

type begin struct {
	path    string
	context string
}

func createBegin(
	path string,
	context string,
) Begin {
	out := begin{
		path:    path,
		context: context,
	}

	return &out
}

// Path returns the path
func (obj *begin) Path() string {
	return obj.path
}

// Context returns the context
func (obj *begin) Context() string {
	return obj.context
}
