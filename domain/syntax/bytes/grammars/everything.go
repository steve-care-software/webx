package grammars

type everything struct {
	name      string
	exception Token
	escape    Token
}

func createEverything(
	name string,
	exception Token,
) Everything {
	return createEverythingInternally(name, exception, nil)
}

func createEverythingWithEscape(
	name string,
	exception Token,
	escape Token,
) Everything {
	return createEverythingInternally(name, exception, escape)
}

func createEverythingInternally(
	name string,
	exception Token,
	escape Token,
) Everything {
	out := everything{
		name:      name,
		exception: exception,
		escape:    escape,
	}

	return &out
}

// Name returns the name
func (obj *everything) Name() string {
	return obj.name
}

// Exception returns the exception
func (obj *everything) Exception() Token {
	return obj.exception
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *everything) HasEscape() bool {
	return obj.escape != nil
}

// Escape returns the escape, if any
func (obj *everything) Escape() Token {
	return obj.escape
}
