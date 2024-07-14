package heads

// NewHeadForTests creates a new head for tests
func NewHeadForTests(context string, ret string) Head {
	ins, err := NewBuilder().Create().WithContext(context).WithReturn(ret).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
