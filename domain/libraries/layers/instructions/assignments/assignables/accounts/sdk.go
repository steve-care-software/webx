package accounts

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/retrieves"
)

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
