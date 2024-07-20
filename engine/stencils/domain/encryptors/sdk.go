package encryptors

// Encryptor represents an encryptor
type Encryptor interface {
	Encrypt(message []byte, password []byte) ([]byte, error)
	Decrypt(cipher []byte, password []byte) ([]byte, error)
}
