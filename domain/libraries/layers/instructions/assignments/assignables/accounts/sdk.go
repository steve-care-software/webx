package accounts

// Account represents an account assignable
type Account interface {
	IsList() bool
	List() string
	IsAccount() bool
	Account() string
	IsCredentials() bool
	Credentials() Credentials
	IsRetrieve() bool
	Retrieve() Retrieve
	IsSign() bool
	Sign() Sign
	IsGenerateRing() bool
	GenerateRing() uint
	IsVote() bool
	Vote() Vote
	IsEncrypt() bool
	Encrypt() Encrypt
	IsDecrypt() bool
	Decrypt() Decrypt
}

// Credentials represents a credentials
type Credentials interface {
	Username() string
	Password() []byte
}

// Retrieve represents a retrieve
type Retrieve interface {
	Password() []byte
	Credentials() string
}

// Sign represenst a sign
type Sign interface {
	Message() string
	Account() string
}

// Vote represents a vote
type Vote interface {
	Message() string
	Ring() string
	Account() string
}

// Encrypt represents an encrypt
type Encrypt interface {
	Message() string
	Account() string
}

// Decrypt represents a decrypt
type Decrypt interface {
	Cipher() string
	Account() string
}
