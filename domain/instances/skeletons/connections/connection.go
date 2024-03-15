package connections

import "github.com/steve-care-software/datastencil/domain/hash"

type connection struct {
	hash hash.Hash
	name string
	from Field
	to   Field
}

func createConnection(
	hash hash.Hash,
	name string,
	from Field,
	to Field,
) Connection {
	out := connection{
		hash: hash,
		name: name,
		from: from,
		to:   to,
	}

	return &out
}

// Hash returns the hash
func (obj *connection) Hash() hash.Hash {
	return obj.hash
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
