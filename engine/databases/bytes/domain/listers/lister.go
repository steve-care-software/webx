package listers

type lister struct {
	keyname string
	index   uint64
	length  uint64
}

func createLister(
	keyname string,
	index uint64,
	length uint64,
) Lister {
	out := lister{
		keyname: keyname,
		index:   index,
		length:  length,
	}

	return &out
}

// Keyname returns the keyname
func (obj *lister) Keyname() string {
	return obj.keyname
}

// Index returns the index
func (obj *lister) Index() uint64 {
	return obj.index
}

// Length returns the length
func (obj *lister) Length() uint64 {
	return obj.length
}
