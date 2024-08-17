package replacements

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithList(list []Replacement) Builder
	Now() (Replacements, error)
}

// Replacements represents a replacement list
type Replacements interface {
	List() []Replacement
}

// ReplacementBuilder represents a replacement builder
type ReplacementBuilder interface {
	Create() ReplacementBuilder
	WithOrigin(origin string) ReplacementBuilder
	WithTarget(target string) ReplacementBuilder
	Now() (Replacement, error)
}

// Replacement represents a replacement
type Replacement interface {
	Origin() string // origin token
	Target() string // target ast
}
