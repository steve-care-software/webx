package databases

// NewDatabaseForTests creates a database for tests
func NewDatabaseForTests(path string, description string, head string, isActive string) Database {
	ins, err := NewBuilder().Create().
		WithPath(path).
		WithDescription(description).
		WithHead(head).
		WithActive(isActive).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
