package encryptions

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/domain/keys/encryptors"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	decryptApp        decrypts.Application
	encryptApp        encrypts.Application
	assignableBuilder stacks.AssignableBuilder
	pkBuilder         encryptors.Builder
	bitRate           int
}

func createApplication(
	decryptApp decrypts.Application,
	encryptApp encrypts.Application,
	assignableBuilder stacks.AssignableBuilder,
	pkBuilder encryptors.Builder,
	bitRate int,
) Application {
	out := application{
		decryptApp:        decryptApp,
		encryptApp:        encryptApp,
		assignableBuilder: assignableBuilder,
		pkBuilder:         pkBuilder,
		bitRate:           bitRate,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable encryptions.Encryption) (stacks.Assignable, *uint, error) {
	if assignable.IsEncrypt() {
		encrypt := assignable.Encrypt()
		return app.encryptApp.Execute(frame, encrypt)
	}

	if assignable.IsDecrypt() {
		decrypt := assignable.Decrypt()
		return app.decryptApp.Execute(frame, decrypt)
	}

	if assignable.IsFetchPublicKey() {
		pkName := assignable.FetchPublicKey()
		pk, err := frame.FetchEncryptor(pkName)
		if err != nil {
			code := failures.CouldNotFetchEncryptionPrivateKeyFromFrame
			return nil, &code, err
		}

		pubKey := pk.Public()
		ins, err := app.assignableBuilder.Create().WithEncryptorPubKey(pubKey).Now()
		if err != nil {
			return nil, nil, err
		}

		if err != nil {
			return nil, nil, err
		}

		return ins, nil, nil
	}

	pPrivateKey, err := rsa.GenerateKey(rand.Reader, app.bitRate)
	if err != nil {
		return nil, nil, err
	}

	encryptor, err := app.pkBuilder.Create().WithPK(*pPrivateKey).WithBitRate(app.bitRate).Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithEncryptor(encryptor).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
