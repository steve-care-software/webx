package publics

type publics struct {
	list []Public
}

func createPublics(
	list []Public,
) Publics {
	out := publics{
		list: list,
	}

	return &out
}

// List returns the publics
func (obj *publics) List() []Public {
	return obj.list
}
