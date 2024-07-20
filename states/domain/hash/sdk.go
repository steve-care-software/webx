package hash

// Size represents the byte size of the hash
const Size = 64

// NewAdapter returns a new hash adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// Hash represents a hash
type Hash []byte

// Adapter represents an hash adapter
type Adapter interface {
	FromBytes(input []byte) (*Hash, error)
	FromMultiBytes(input [][]byte) (*Hash, error)
	FromString(input string) (*Hash, error)
}
