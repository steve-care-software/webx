package cryptography

// Cryptography represents a cryptography
type Cryptography interface {
	IsHash() bool
	Hash() string
	IsEncrypt() bool
	Encrypt() Encrypt
	IsDecrypt() bool
	Decrypt() Decrypt
}

// Encrypt represents an encrypt
type Encrypt interface {
	Message() string
	Password() string
}

// Decrypt represents a decrypt
type Decrypt interface {
	Cipher() string
	Password() string
}
