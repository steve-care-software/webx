package connections

type connection struct {
	name string
	from Field
	to   Field
}

func createConnection(
	name string,
	from Field,
	to Field,
) Connection {
	out := connection{
		name: name,
		from: from,
		to:   to,
	}

	return &out
}

// Name returns the name
func (obj *connection) Name() string {
	return obj.name
}

// From returns the from
func (obj *connection) From() Field {
	return obj.from
}

// To returns the to
func (obj *connection) To() Field {
	return obj.to
}
