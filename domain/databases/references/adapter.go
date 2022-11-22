package references

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToContent converts reference to bytes
func (app *adapter) ToContent(ins Reference) ([]byte, error) {
	return nil, nil
}

// ToReference converts bytes to reference
func (app *adapter) ToReference(content []byte) (Reference, error) {
	return nil, nil
}
