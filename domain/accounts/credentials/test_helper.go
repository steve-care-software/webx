package credentials

// NewCredentialsForTests creates a new credentials for tests
func NewCredentialsForTests(username string, password []byte) Credentials {
	ins, err := NewBuilder().Create().
		WithUsername(username).
		WithPassword(password).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
