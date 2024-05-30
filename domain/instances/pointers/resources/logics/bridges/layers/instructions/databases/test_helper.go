package databases

// NewDatabaseWithSaveForTests creates a new database with save for tests
func NewDatabaseWithSaveForTests(save string) Database {
	ins, err := NewBuilder().Create().WithSave(save).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithDeleteForTests creates a new database with delete for tests
func NewDatabaseWithDeleteForTests(delete string) Database {
	ins, err := NewBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
