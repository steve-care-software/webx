package encryptions

// Application represents an encryption application
type Application interface {
	Encrypt(message []byte, password []byte) ([]byte, error)
	Decrypt(cipher []byte, password []byte) ([]byte, error)
}
