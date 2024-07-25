package omissions

import "github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"

// NewOmissionWithPrefixAndSuffixForTests creates omission with prefix and suffix for tests
func NewOmissionWithPrefixAndSuffixForTests(prefix elements.Element, suffix elements.Element) Omission {
	ins, err := NewBuilder().Create().WithPrefix(prefix).WithSuffix(suffix).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOmissionWithPrefixForTests creates omission with prefix for tests
func NewOmissionWithPrefixForTests(prefix elements.Element) Omission {
	ins, err := NewBuilder().Create().WithPrefix(prefix).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOmissionWithSuffixForTests creates omission with suffix for tests
func NewOmissionWithSuffixForTests(suffix elements.Element) Omission {
	ins, err := NewBuilder().Create().WithSuffix(suffix).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
