package encrypts

// Encrypt represents an encrypt
type Encrypt struct {
	Message   string `json:"message"`
	PublicKey string `json:"pubkey"`
}
