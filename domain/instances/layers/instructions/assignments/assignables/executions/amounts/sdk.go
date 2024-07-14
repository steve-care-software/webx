package amounts

// Amount represents an amount
type Amount interface {
	Context() string
	Return() string
}
