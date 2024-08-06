package keys

// Key represents the identity keys
type Key interface {
	Signer() []byte
	Encryptor() []byte
}
