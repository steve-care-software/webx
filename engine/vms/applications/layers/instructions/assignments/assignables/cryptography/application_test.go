package cryptography

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/domain/signers"
	application_decrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	application_encrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	application_keys "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys"
	application_encryptions "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	application_keys_decrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	application_keys_encrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	application_signatures "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	application_signs "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	application_signs_creates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	application_signs_validates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	application_votes "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	application_votes_creates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	application_votes_validates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/mocks"
)

func TestExecute_withDecrypt_Success(t *testing.T) {
	cipherVar := "cipherVar"
	cipher := []byte("this is a cipher")
	passwordVar := "passVar"
	password := []byte("this is a password")
	message := []byte("this is some message")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(
					cipher,
				),
			),
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					password,
				),
			),
		}),
	)

	instruction := cryptography.NewCryptographyWithDecryptForTests(
		decrypts.NewDecryptForTests(cipherVar, passwordVar),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{
			string(cipher): map[string][]byte{
				string(password): message,
			},
		},
	)

	application := NewApplication(
		application_decrypts.NewApplication(
			encryptor,
		),
		application_encrypts.NewApplication(
			encryptor,
		),
		application_keys.NewApplication(
			application_encryptions.NewApplication(
				application_keys_decrypts.NewApplication(),
				application_keys_encrypts.NewApplication(),
				1048,
			),
			application_signatures.NewApplication(
				application_votes.NewApplication(
					application_votes_creates.NewApplication(),
					application_votes_validates.NewApplication(),
				),
				application_signs.NewApplication(
					application_signs_creates.NewApplication(),
					application_signs_validates.NewApplication(),
				),
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retMessage := retAssignable.Bytes()
	if !bytes.Equal(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestExecute_withEncrypt_Success(t *testing.T) {
	messageVar := "messageVar"
	message := []byte("this is a message")
	passwordVar := "passVar"
	password := []byte("this is a password")
	cipher := []byte("this is some cipher")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(
					message,
				),
			),
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					password,
				),
			),
		}),
	)

	instruction := cryptography.NewCryptographyWithEncryptForTests(
		encrypts.NewEncryptForTests(messageVar, passwordVar),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{
			string(message): map[string][]byte{
				string(password): cipher,
			},
		},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_decrypts.NewApplication(
			encryptor,
		),
		application_encrypts.NewApplication(
			encryptor,
		),
		application_keys.NewApplication(
			application_encryptions.NewApplication(
				application_keys_decrypts.NewApplication(),
				application_keys_encrypts.NewApplication(),
				1048,
			),
			application_signatures.NewApplication(
				application_votes.NewApplication(
					application_votes_creates.NewApplication(),
					application_votes_validates.NewApplication(),
				),
				application_signs.NewApplication(
					application_signs_creates.NewApplication(),
					application_signs_validates.NewApplication(),
				),
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retCipher := retAssignable.Bytes()
	if !bytes.Equal(cipher, retCipher) {
		t.Errorf("the returned cipher is invalid")
		return
	}
}

func TestExecute_withKey_Success(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithSignerPublicKeyForTests(
			firstPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			secondPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			thirdPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			signer.PublicKey(),
		),
	})

	messageVar := "myMessage"
	ringVar := "myVar"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := cryptography.NewCryptographyWithKeyForTests(
		keys.NewKeyWithSignatureForTests(
			signatures.NewSignatureWithGeneratePrivateKeyForTests(),
		),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_decrypts.NewApplication(
			encryptor,
		),
		application_encrypts.NewApplication(
			encryptor,
		),
		application_keys.NewApplication(
			application_encryptions.NewApplication(
				application_keys_decrypts.NewApplication(),
				application_keys_encrypts.NewApplication(),
				1048,
			),
			application_signatures.NewApplication(
				application_votes.NewApplication(
					application_votes_creates.NewApplication(),
					application_votes_validates.NewApplication(),
				),
				application_signs.NewApplication(
					application_signs_creates.NewApplication(),
					application_signs_validates.NewApplication(),
				),
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsSigner() {
		t.Errorf("the assignable was expected to contain a signer")
		return
	}
}
