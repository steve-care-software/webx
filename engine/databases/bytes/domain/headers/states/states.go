package states

type states struct {
	list []State
}

func createStates(
	list []State,
) States {
	out := states{
		list: list,
	}

	return &out
}

// List returns the list of states
func (obj *states) List() []State {
	return obj.list
}
