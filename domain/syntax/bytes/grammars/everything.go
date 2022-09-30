package grammars

type everything struct {
	exception Line
	escape    Line
}

func createEverything(
	exception Line,
) Everything {
	return createEverythingInternally(exception, nil)
}

func createEverythingWithEscape(
	exception Line,
	escape Line,
) Everything {
	return createEverythingInternally(exception, escape)
}

func createEverythingInternally(
	exception Line,
	escape Line,
) Everything {
	out := everything{
		exception: exception,
		escape:    escape,
	}

	return &out
}

// Exception returns the exception
func (obj *everything) Exception() Line {
	return obj.exception
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *everything) HasEscape() bool {
	return obj.escape != nil
}

// Escape returns the escape, if any
func (obj *everything) Escape() Line {
	return obj.escape
}
