package accounts

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/retrieves"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithList(list string) Builder
	WithCredentials(credentials credentials.Credentials) Builder
	WithRetrieve(retrieve retrieves.Retrieve) Builder
	WithCommunication(communication communications.Communication) Builder
	WithEncryption(encryption encryptions.Encryption) Builder
	Now() (Account, error)
}

// Account represents an account assignable
type Account interface {
	Hash() hash.Hash
	IsList() bool
	List() string
	IsCredentials() bool
	Credentials() credentials.Credentials
	IsRetrieve() bool
	Retrieve() retrieves.Retrieve
	IsCommunication() bool
	Communication() communications.Communication
	IsEncryption() bool
	Encryption() encryptions.Encryption
}
