package files

// NewFileWithCloseForTests creates a new file with close
func NewFileWithCloseForTests(close string) File {
	ins, err := NewBuilder().Create().WithClose(close).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFileWithDeleteForTests creates a new file with delete
func NewFileWithDeleteForTests(delete string) File {
	ins, err := NewBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
