package rules

// NewRulesForTests creates a new rules for tests
func NewRulesForTests(list []Rule) Rules {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRuleForTests creates a new rule for tests
func NewRuleForTests(name string, values []byte) Rule {
	ins, err := NewRuleBuilder().Create().WithName(name).WithBytes(values).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
