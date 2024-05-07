package transforms

// NewTransformForTests creates a new transform for tests
func NewTransformForTests(query []byte, data []byte) Transform {
	ins, err := NewBuilder().Create().WithQuery(query).WithBytes(data).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
