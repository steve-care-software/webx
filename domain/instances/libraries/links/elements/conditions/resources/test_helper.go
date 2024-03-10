package resources

// NewResourceWithIsRaisedInLayerForTests creates a new condition resource withIsRaisedInLayer for tests
func NewResourceWithIsRaisedInLayerForTests(code uint) Resource {
	ins, err := NewBuilder().Create().WithCode(code).IsRaisedInLayer().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResourceForTests creates a new condition resource for tests
func NewResourceForTests(code uint) Resource {
	ins, err := NewBuilder().Create().WithCode(code).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
