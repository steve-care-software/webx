package retrievals

type retrieval struct {
	index  uint64
	length uint64
}

func createRetrieval(
	index uint64,
	length uint64,
) Retrieval {
	out := retrieval{
		index:  index,
		length: length,
	}

	return &out
}

// Index returns the index
func (obj *retrieval) Index() uint64 {
	return obj.index
}

// Length returns the length
func (obj *retrieval) Length() uint64 {
	return obj.length
}
