package retrieves

// NewRetrieveForTests creates a new retrieve for tests
func NewRetrieveForTests(context string, index string, ret string) Retrieve {
	ins, err := NewBuilder().Create().WithContext(context).WithIndex(index).WithReturn(ret).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRetrieveWithLengthForTests creates a new retrieve with length for tests
func NewRetrieveWithLengthForTests(context string, index string, ret string, length string) Retrieve {
	ins, err := NewBuilder().Create().WithContext(context).WithIndex(index).WithReturn(ret).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
