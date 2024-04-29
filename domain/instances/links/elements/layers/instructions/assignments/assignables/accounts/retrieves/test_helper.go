package retrieves

// NewRetrieveForTests creates a new retrieve for tests
func NewRetrieveForTests(password string, credentials string) Retrieve {
	ins, err := NewBuilder().Create().WithPassword(password).WithCredentials(credentials).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
