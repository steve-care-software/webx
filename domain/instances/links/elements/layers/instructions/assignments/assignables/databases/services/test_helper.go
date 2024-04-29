package services

// NewServiceWithBeginForTests creates a new service with begin for tests
func NewServiceWithBeginForTests() Service {
	ins, err := NewBuilder().Create().IsBegin().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
