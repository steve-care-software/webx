package cardinalities

// Cardinality represents a cardinality
type Cardinality struct {
	Min uint  `json:"min"`
	Max *uint `json:"max"`
}
