package cryptography

// Application represents a cryptography application
type Application interface {
	Encrypt(message []byte, password byte) ([]byte, error)  // encrypt data using a password
	Decrypt(cipher []byte, password []byte) ([]byte, error) // decrypt a cipher using a password
	GeneratePrivateKey(words []string)                      // generate a private key and returns it
	Sign(message []byte, pk string)                         // sign a message using a private key
	VerifySignature(signature string, data []byte)          // verify a signature
}
