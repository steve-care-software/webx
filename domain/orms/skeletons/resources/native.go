package resources

type native struct {
	pSingle *uint8
	list    List
}

func createNativeWithSingle(
	pSingle *uint8,
) Native {
	return createNativeInternally(pSingle, nil)
}

func createNativeWithList(
	list List,
) Native {
	return createNativeInternally(nil, list)
}

func createNativeInternally(
	pSingle *uint8,
	list List,
) Native {
	out := native{
		pSingle: pSingle,
		list:    list,
	}

	return &out
}

// IsSingle returns true if there is a single, flase otherwise
func (obj *native) IsSingle() bool {
	return obj.pSingle != nil
}

// Single returns the single, if any
func (obj *native) Single() *uint8 {
	return obj.pSingle
}

// IsList returns true if there is a list, flase otherwise
func (obj *native) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *native) List() List {
	return obj.list
}
