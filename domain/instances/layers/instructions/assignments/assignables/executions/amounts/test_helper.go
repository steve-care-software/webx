package amounts

// NewAmountForTests creates a new amount for tests
func NewAmountForTests(context string, ret string) Amount {
	ins, err := NewBuilder().Create().WithContext(context).WithReturn(ret).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
