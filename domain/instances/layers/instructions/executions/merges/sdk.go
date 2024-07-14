package merges

// Merge represents a merge
type Merge interface {
	Base() string
	Top() string
}
