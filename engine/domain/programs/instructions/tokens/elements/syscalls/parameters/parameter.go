package parameters

type parameter struct {
	token string
	index uint
	name  string
}

func createParameter(
	token string,
	index uint,
	name string,
) Parameter {
	out := parameter{
		token: token,
		index: index,
		name:  name,
	}

	return &out
}

// Token returns the token
func (obj *parameter) Token() string {
	return obj.token
}

// Index returns the index
func (obj *parameter) Index() uint {
	return obj.index
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}
