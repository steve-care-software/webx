package binaries

// Application represents a binary application
type Application interface {
	Execute(binaries []byte, cmd []string) ([]byte, error)
}
