package decrypts

// Decrypt represents a decrypt
type Decrypt struct {
	Cipher     string `json:"cipher"`
	PrivateKey string `json:"pk"`
}
