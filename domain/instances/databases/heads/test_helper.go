package heads

// NewHeadForTests creates a new head for tests
func NewHeadForTests(path []string, description string, isActive bool) Head {
	builder := NewBuilder().Create().WithPath(path).WithDescription(description)
	if isActive {
		builder.IsActive()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
