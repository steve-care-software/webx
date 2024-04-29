package accounts

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/retrieves"
)

type account struct {
	hash          hash.Hash
	list          string
	credentials   credentials.Credentials
	retrieve      retrieves.Retrieve
	communication communications.Communication
	encryption    encryptions.Encryption
}

func createAccountWithList(
	hash hash.Hash,
	list string,
) Account {
	return createAccountInternally(hash, list, nil, nil, nil, nil)
}

func createAccountWithCredentials(
	hash hash.Hash,
	credentials credentials.Credentials,
) Account {
	return createAccountInternally(hash, "", credentials, nil, nil, nil)
}

func createAccountWithRetrieve(
	hash hash.Hash,
	retrieve retrieves.Retrieve,
) Account {
	return createAccountInternally(hash, "", nil, retrieve, nil, nil)
}

func createAccountWithCommunication(
	hash hash.Hash,
	communication communications.Communication,
) Account {
	return createAccountInternally(hash, "", nil, nil, communication, nil)
}

func createAccountWithEncryption(
	hash hash.Hash,
	encryption encryptions.Encryption,
) Account {
	return createAccountInternally(hash, "", nil, nil, nil, encryption)
}

func createAccountInternally(
	hash hash.Hash,
	list string,
	credentials credentials.Credentials,
	retrieve retrieves.Retrieve,
	communication communications.Communication,
	encryption encryptions.Encryption,
) Account {
	out := account{
		hash:          hash,
		list:          list,
		credentials:   credentials,
		retrieve:      retrieve,
		communication: communication,
		encryption:    encryption,
	}

	return &out
}

// Hash returns the hash
func (obj *account) Hash() hash.Hash {
	return obj.hash
}

// IsList returns true if there is a list, false otherwise
func (obj *account) IsList() bool {
	return obj.list != ""
}

// List returns the list, if any
func (obj *account) List() string {
	return obj.list
}

// IsCredentials returns true if there is a list, false otherwise
func (obj *account) IsCredentials() bool {
	return obj.credentials != nil
}

// Credentials returns the list, if any
func (obj *account) Credentials() credentials.Credentials {
	return obj.credentials
}

// IsCredentials returns true if there is a list, false otherwise
func (obj *account) IsRetrieve() bool {
	return obj.retrieve != nil
}

// Retrieve returns the retrieve, if any
func (obj *account) Retrieve() retrieves.Retrieve {
	return obj.retrieve
}

// IsCommunication returns true if there is a communication, false otherwise
func (obj *account) IsCommunication() bool {
	return obj.communication != nil
}

// Communication returns the communication, if any
func (obj *account) Communication() communications.Communication {
	return obj.communication
}

// IsEncryption returns true if there is an encryption, false otherwise
func (obj *account) IsEncryption() bool {
	return obj.encryption != nil
}

// Encryption returns the encryption, if any
func (obj *account) Encryption() encryptions.Encryption {
	return obj.encryption
}
