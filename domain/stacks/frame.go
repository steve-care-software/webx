package stacks

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/keys/encryptors"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
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

// FetchStringList fetches a string list by name
func (obj *frame) FetchStringList(name string) ([]string, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsStringList() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a []string", name)
		return nil, errors.New(str)
	}

	return assignable.StringList(), nil
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
	return nil, nil
}

// FetchEncryptorPubKey fetches an encryptor pubKey by name
func (obj *frame) FetchEncryptorPubKey(name string) (encryptors.PublicKey, error) {
	return nil, nil
}

// FetchSigner fetches a signer pk by name
func (obj *frame) FetchSigner(name string) (signers.Signer, error) {
	return nil, nil
}

// FetchSignerPubKey fetches a signer pubKey by name
func (obj *frame) FetchSignerPubKey(name string) (signers.PublicKey, error) {
	return nil, nil
}

// FetchSignature fetches a signature by name
func (obj *frame) FetchSignature(name string) (signers.Signature, error) {
	return nil, nil
}

// FetchVote fetches a vote by name
func (obj *frame) FetchVote(name string) (signers.Vote, error) {
	return nil, nil
}

// FetchRing fetches a ring by name
func (obj *frame) FetchRing(name string) ([]signers.PublicKey, error) {
	return nil, nil
}

// FetchList fetches a list by name
func (obj *frame) FetchList(name string) (Assignables, error) {
	return nil, nil
}

// FetchModifications fetches a modifications by name
func (obj *frame) FetchModifications(name string) (modifications.Modifications, error) {
	return nil, nil
}

// FetchCommit fetches a commit by name
func (obj *frame) FetchCommit(name string) (commits.Commit, error) {
	return nil, nil
}

// FetchString fetches a string by name
func (obj *frame) FetchString(name string) (string, error) {
	return "", nil
}

// FetchDelete fetches a delete by name
func (obj *frame) FetchDelete(name string) (deletes.Delete, error) {
	return nil, nil
}

// HasAssignments returns true if there is assignments, false otherwise
func (obj *frame) HasAssignments() bool {
	return obj.assignments != nil
}

// Assignments fetches the assignments, if any
func (obj *frame) Assignments() Assignments {
	return obj.assignments
}
