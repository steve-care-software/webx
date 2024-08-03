package files

// NewFileWithCloseForTests creates a new file with close
func NewFileWithCloseForTests(close string) File {
	ins, err := NewBuilder().Create().WithClose(close).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
