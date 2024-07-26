package listers

// Lister represents a lister
type Lister interface {
	Keyname() string
	Index() uint64
	Length() uint64
}
