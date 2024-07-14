package heads

// Head represents a head
type Head interface {
	Context() string
	Hash() string
}
