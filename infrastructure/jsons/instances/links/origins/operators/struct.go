package operators

// Operator represents an operator
type Operator struct {
	And bool `json:"and"`
	Or  bool `json:"or"`
	Xor bool `json:"xor"`
}
