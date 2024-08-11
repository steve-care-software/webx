package switchers

type switchersIns struct {
	list []Switcher
}

func createSwitchers(
	list []Switcher,
) Switchers {
	out := switchersIns{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *switchersIns) List() []Switcher {
	return obj.list
}
