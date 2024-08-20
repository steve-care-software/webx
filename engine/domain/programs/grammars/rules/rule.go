package rules

type rule struct {
	name  string
	bytes []byte
}

func createRule(
	name string,
	bytes []byte,
) Rule {
	out := rule{
		name:  name,
		bytes: bytes,
	}

	return &out
}

// Name returns the name
func (obj *rule) Name() string {
	return obj.name
}

// Bytes returns the bytes
func (obj *rule) Bytes() []byte {
	return obj.bytes
}
