package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
)

// ToIdentity creates a new identity instance from json
func ToIdentity(data []byte) (identities.Identity, error) {
	ptr := new(Identity)
	err := json.Unmarshal(data, ptr)
	if err != nil {
		return nil, err
	}

	list := []modifications.Modification{}
	encPKAdapter := keys.NewAdapter()
	sigPKAdapter := signatures.NewPrivateKeyAdapter()
	modificationBuilder := modifications.NewModificationBuilder()
	for _, oneModification := range ptr.Modifications {
		encPK, err := encPKAdapter.FromBytes(oneModification.EncPK)
		if err != nil {
			return nil, err
		}

		sigPK, err := sigPKAdapter.ToPrivateKey(oneModification.SigPK)
		if err != nil {
			return nil, err
		}

		ins, err := modificationBuilder.Create().
			WithName(oneModification.Name).
			WithSignature(sigPK).
			WithEncryption(encPK).
			CreatedOn(oneModification.CreatedOn).
			Now()

		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	modifications, err := modifications.NewBuilder().Create().WithList(list).Now()
	if err != nil {
		return nil, err
	}

	return identities.NewBuilder().Create().
		WithModifications(modifications).
		Now()
}

// ToJSON creates json data from an identity instance
func ToJSON(identity identities.Identity) ([]byte, error) {
	encPKAdapter := keys.NewAdapter()
	list := identity.Modifications().List()
	modifications := []Modification{}
	for _, oneModification := range list {
		content := oneModification.Content()
		ins := Modification{
			Name:      content.Name(),
			SigPK:     content.Signature().String(),
			EncPK:     encPKAdapter.ToBytes(content.Encryption()),
			CreatedOn: oneModification.CreatedOn(),
		}

		modifications = append(modifications, ins)
	}

	ins := Identity{
		Modifications: modifications,
	}

	return json.Marshal(ins)
}
