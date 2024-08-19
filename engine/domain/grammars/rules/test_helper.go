package rules

// NewRuleForTests creates a new rule for tests
func NewRuleForTests(name string, values []byte) Rule {
	ins, err := NewRuleBuilder().Create().WithName(name).WithBytes(values).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
