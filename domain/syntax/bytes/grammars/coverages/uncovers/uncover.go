package uncovers

type uncover struct {
	name string
	line Line
}

func createUncover(
	name string,
	line Line,
) Uncover {
	out := uncover{
		name: name,
		line: line,
	}

	return &out
}

// Name retruns the name
func (obj *uncover) Name() string {
	return obj.name
}

// Line retruns the line
func (obj *uncover) Line() Line {
	return obj.line
}
