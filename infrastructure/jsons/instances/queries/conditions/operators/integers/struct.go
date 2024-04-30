package integers

// Integer represents an integer
type Integer struct {
	IsSmalerThan bool `json:"is_smaller_than"`
	IsBiggerThan bool `json:"is_bigger_than"`
	IsEqual      bool `json:"is_equal"`
}
