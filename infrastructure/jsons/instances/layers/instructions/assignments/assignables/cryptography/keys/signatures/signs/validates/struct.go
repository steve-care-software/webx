package validates

// Validate represents the validate
type Validate struct {
	Signature string `json:"signature"`
	Message   string `json:"message"`
	PublicKey string `json:"pubkey"`
}
