package validates

// Validate represents a validate
type Validate struct {
	Vote       string `json:"vote"`
	Message    string `json:"message"`
	HashedRing string `json:"hashed_ring"`
}
