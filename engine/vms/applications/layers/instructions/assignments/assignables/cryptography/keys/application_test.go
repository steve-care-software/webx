package keys

import (
	"bytes"
	"testing"

	application_encryptions "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	application_decrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	application_encrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	application_signatures "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	application_signs "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	application_signs_creates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	application_signs_validates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	application_votes "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	application_votes_creates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	application_votes_validates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/encryptors"
	"github.com/steve-care-software/webx/engine/vms/domain/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

func TestExecute_withSignature_Success(t *testing.T) {
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

	instruction := keys.NewKeyWithSignatureForTests(
		signatures.NewSignatureWithGeneratePrivateKeyForTests(),
	)

	application := NewApplication(
		application_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
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

func TestExecute_withEncryption_Success(t *testing.T) {
	message := []byte("this is some data")
	encryptor := encryptors.NewEncryptorForTests(1048)
	pubKey := encryptor.Public()

	messageVar := "myMessage"
	pubKeyVar := "myCipher"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				pubKeyVar,
				stacks.NewAssignableWithEncryptorPublicKeyForTests(pubKey),
			),
		}),
	)

	instruction := keys.NewKeyWithEncryptionForTests(
		encryptions.NewEncryptionWithEncryptForTests(
			encrypts.NewEncryptForTests(messageVar, pubKeyVar),
		),
	)

	application := NewApplication(
		application_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
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
	retMessage, err := encryptor.Decrypt(retCipher)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}
