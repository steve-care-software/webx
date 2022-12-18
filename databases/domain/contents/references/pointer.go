package references

type pointer struct {
	from   uint
	length uint
}

func createPointer(
	from uint,
	length uint,
) Pointer {
	out := pointer{
		from:   from,
		length: length,
	}

	return &out
}

// From retruns the from
func (obj *pointer) From() uint {
	return obj.from
}

// Length retruns the length
func (obj *pointer) Length() uint {
	return obj.length
}
