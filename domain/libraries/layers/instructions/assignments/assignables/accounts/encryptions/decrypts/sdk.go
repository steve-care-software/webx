package decrypts

// Builder represents a decrypt builder
type Builder interface {
	Create() Builder
	WithCipher(cipher string) Builder
	WithAmount(amount string) Builder
	Now() (Decrypt, error)
}

// Decrypt represents a decrypt
type Decrypt interface {
	Cipher() string
	Account() string
}
