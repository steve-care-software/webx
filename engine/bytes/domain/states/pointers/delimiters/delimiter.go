package delimiters

type delimiter struct {
	index  uint64
	length uint64
}

func createDelimiter(
	index uint64,
	length uint64,
) Delimiter {
	out := delimiter{
		index:  index,
		length: length,
	}

	return &out
}

// Index returns the index
func (obj *delimiter) Index() uint64 {
	return obj.index
}

// Length returns the length
func (obj *delimiter) Length() uint64 {
	return obj.length
}
