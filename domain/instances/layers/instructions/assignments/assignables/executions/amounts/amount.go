package amounts

type amount struct {
	context string
	ret     string
}

func createAmount(
	context string,
	ret string,
) Amount {
	out := amount{
		context: context,
		ret:     ret,
	}

	return &out
}

// Context returns the context
func (obj *amount) Context() string {
	return obj.context
}

// Return returns the retunr
func (obj *amount) Return() string {
	return obj.ret
}
