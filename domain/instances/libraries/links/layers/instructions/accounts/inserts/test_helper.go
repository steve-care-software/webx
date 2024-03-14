package inserts

// NewInsertForTests creates a new insert for tests
func NewInsertForTests(username string, password string) Insert {
	ins, err := NewBuilder().Create().
		WithUsername(username).
		WithPassword(password).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
