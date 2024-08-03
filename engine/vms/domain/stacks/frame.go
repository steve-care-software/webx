package stacks

import (
	"errors"
	"fmt"
	"os"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/units/domain/identities/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/encryptors"
)

type frame struct {
	assignments Assignments
}

func createFrame() Frame {
	return createFrameInternally(
		nil,
	)
}

func createFrameWithAssignments(
	assignments Assignments,
) Frame {
	return createFrameInternally(
		assignments,
	)
}

func createFrameInternally(
	assignments Assignments,
) Frame {
	out := frame{
		assignments: assignments,
	}

	return &out
}

// Fetch fetches an assignable by name
func (obj *frame) Fetch(name string) (Assignable, error) {
	if !obj.HasAssignments() {
		str := fmt.Sprintf("the frame contains no assignment, therefore the assignable (name: %s) could not be found", name)
		return nil, errors.New(str)
	}

	return obj.Assignments().Fetch(name)
}

// FetchBool fetches a bool by name
func (obj *frame) FetchBool(name string) (bool, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return false, err
	}

	if !assignable.IsBool() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a bool", name)
		return false, errors.New(str)
	}

	pBool := assignable.Bool()
	return *pBool, nil
}

// FetchHash fetches an hash by name
func (obj *frame) FetchHash(name string) (hash.Hash, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsHash() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Hash", name)
		return nil, errors.New(str)
	}

	return assignable.Hash(), nil
}

// FetchBytes fetches a bytes by name
func (obj *frame) FetchBytes(name string) ([]byte, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsBytes() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a []byte", name)
		return nil, errors.New(str)
	}

	return assignable.Bytes(), nil
}

// FetchUnsignedInt fetches an unsigned int by name
func (obj *frame) FetchUnsignedInt(name string) (*uint, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsUnsignedInt() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a uint", name)
		return nil, errors.New(str)
	}

	return assignable.UnsignedInt(), nil
}

// FetchInstance fetches an instance by name
func (obj *frame) FetchInstance(name string) (instances.Instance, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsInstance() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Instance", name)
		return nil, errors.New(str)
	}

	return assignable.Instance(), nil
}

// FetchEncryptor fetches an encryptor by name
func (obj *frame) FetchEncryptor(name string) (encryptors.Encryptor, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsEncryptor() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Encryptor", name)
		return nil, errors.New(str)
	}

	return assignable.Encryptor(), nil
}

// FetchEncryptorPubKey fetches an encryptor pubKey by name
func (obj *frame) FetchEncryptorPubKey(name string) (encryptors.PublicKey, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsEncryptorPublicKey() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Encryptor's PublicKey", name)
		return nil, errors.New(str)
	}

	return assignable.EncryptorPublicKey(), nil
}

// FetchSigner fetches a signer pk by name
func (obj *frame) FetchSigner(name string) (signers.Signer, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsSigner() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a Signer", name)
		return nil, errors.New(str)
	}

	return assignable.Signer(), nil
}

// FetchSignerPubKey fetches a signer pubKey by name
func (obj *frame) FetchSignerPubKey(name string) (signers.PublicKey, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsSignerPublicKey() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a Signer's PublicKey", name)
		return nil, errors.New(str)
	}

	return assignable.SignerPublicKey(), nil
}

// FetchSignature fetches a signature by name
func (obj *frame) FetchSignature(name string) (signers.Signature, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsSignature() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a Signature", name)
		return nil, errors.New(str)
	}

	return assignable.Signature(), nil
}

// FetchVote fetches a vote by name
func (obj *frame) FetchVote(name string) (signers.Vote, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsVote() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a Vote", name)
		return nil, errors.New(str)
	}

	return assignable.Vote(), nil
}

// FetchRing fetches a ring by name
func (obj *frame) FetchRing(name string) ([]signers.PublicKey, error) {
	assignables, err := obj.FetchList(name)
	if err != nil {
		return nil, err
	}

	list := assignables.List()
	output := []signers.PublicKey{}
	for _, oneAssignable := range list {
		if !oneAssignable.IsSignerPublicKey() {
			str := fmt.Sprintf("the assignable (name: %s) was expected to contain a Signer's Public Key as its elements", name)
			return nil, errors.New(str)
		}

		output = append(output, oneAssignable.SignerPublicKey())
	}

	return output, nil
}

// FetchList fetches a list by name
func (obj *frame) FetchList(name string) (Assignables, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsList() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a List", name)
		return nil, errors.New(str)
	}

	return assignable.List(), nil
}

// FetchString fetches a string by name
func (obj *frame) FetchString(name string) (string, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return "", err
	}

	if !assignable.IsString() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a String", name)
		return "", errors.New(str)
	}

	pString := assignable.String()
	return *pString, nil
}

// FetchApplication fetches an application by name
func (obj *frame) FetchApplication(name string) (applications.Application, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsApplication() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Application", name)
		return nil, errors.New(str)
	}

	retApp := assignable.Applicattion()
	return retApp, nil
}

// FetchFile fetches file identifier by name
func (obj *frame) FetchFile(name string) (*os.File, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsFilePointer() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a file pointer", name)
		return nil, errors.New(str)
	}

	retFile := assignable.FilePointer()
	return retFile, nil
}

// HasAssignments returns true if there is assignments, false otherwise
func (obj *frame) HasAssignments() bool {
	return obj.assignments != nil
}

// Assignments fetches the assignments, if any
func (obj *frame) Assignments() Assignments {
	return obj.assignments
}
