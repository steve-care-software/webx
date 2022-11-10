package heads

type link struct {
	from uint
	to   uint
}

func createLink(
	from uint,
	to uint,
) Link {
	out := link{
		from: from,
		to:   to,
	}

	return &out
}

// From returns the from index
func (obj *link) From() uint {
	return obj.from
}

// To returns the to index
func (obj *link) To() uint {
	return obj.to
}
