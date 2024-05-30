package contents

// Repository represents a content repository
type Repository interface {
	Retrieve(path []string) ([]byte, error)
}
