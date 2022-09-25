package compilers

type composition struct {
	prefix       []byte
	suffix       []byte
	pattern      []byte
	replacements Replacements
}

func createComposition(
	prefix []byte,
	suffix []byte,
	pattern []byte,
	replacements Replacements,
) Composition {
	out := composition{
		prefix:       prefix,
		suffix:       suffix,
		pattern:      pattern,
		replacements: replacements,
	}

	return &out
}

// Prefix returns the prefix
func (obj *composition) Prefix() []byte {
	return obj.prefix
}

// Suffix returns the suffix
func (obj *composition) Suffix() []byte {
	return obj.suffix
}

// Pattern returns the pattern
func (obj *composition) Pattern() []byte {
	return obj.pattern
}

// Replacements returns the replacements
func (obj *composition) Replacements() Replacements {
	return obj.replacements
}
