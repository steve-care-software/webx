package databases

// Database represents a database instruction
type Database interface {
	IsInsert() bool
	Insert() string
	IsDelete() bool
	Delete() string
}
