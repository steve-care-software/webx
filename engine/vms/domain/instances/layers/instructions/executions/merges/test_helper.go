package merges

// NewMergeForTests creates a new merge for tests
func NewMergeForTests(base string, top string) Merge {
	ins, err := NewBuilder().Create().WithBase(base).WithTop(top).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
