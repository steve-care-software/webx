package databases

type migration struct {
	previous    Head
	height      uint
	description string
}

func createMigration(
	previous Head,
	height uint,
	description string,
) Migration {
	out := migration{
		previous:    previous,
		height:      height,
		description: description,
	}

	return &out
}

// Previous returns the previous head
func (obj *migration) Previous() Head {
	return obj.previous
}

// Height returns the height
func (obj *migration) Height() uint {
	return obj.height
}

// Description returns the description
func (obj *migration) Description() string {
	return obj.description
}
