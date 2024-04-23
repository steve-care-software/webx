package criterias

// Criteria represents a criteria
type Criteria struct {
	ChangeSigner    bool   `json:"change_signer"`
	ChangeEncryptor bool   `json:"change_encryptor"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}
