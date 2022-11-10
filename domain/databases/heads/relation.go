package heads

type relation struct {
	from uint
	to   []uint
}

func createRelation(
	from uint,
	to []uint,
) Relation {
	out := relation{
		from: from,
		to:   to,
	}

	return &out
}

// From returns the from index
func (obj *relation) From() uint {
	return obj.from
}

// To returns the to indexes
func (obj *relation) To() []uint {
	return obj.to
}
