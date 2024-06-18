package operators

// Operator represents an operator
type Operator struct {
	IsAnd bool `json:"and"`
	IsOr  bool `json:"or"`
	IsXor bool `json:"xor"`
}
