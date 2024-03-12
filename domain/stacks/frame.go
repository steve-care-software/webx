package stacks

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/queries"
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

// FetchAccount fetches an account by name
func (obj *frame) FetchAccount(name string) (accounts.Account, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	str := "the assignable (name: %s) was expected to contain an Account"
	if !assignable.IsAccount() {
		str := fmt.Sprintf(str, name)
		return nil, errors.New(str)
	}

	account := assignable.Account()
	if !account.IsAccount() {
		str := fmt.Sprintf(str, name)
		return nil, errors.New(str)
	}

	return account.Account(), nil
}

// FetchRing fetches a ring by name
func (obj *frame) FetchRing(name string) ([]signers.PublicKey, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsAccount() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Account", name)
		return nil, errors.New(str)
	}

	account := assignable.Account()
	if !account.IsRing() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a PublicKey ring", name)
		return nil, errors.New(str)
	}

	return account.Ring(), nil
}

// FetchCredentials fetches a credentials by name
func (obj *frame) FetchCredentials(name string) (credentials.Credentials, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsAccount() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain an Account", name)
		return nil, errors.New(str)
	}

	account := assignable.Account()
	if !account.IsCredentials() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a Credentials", name)
		return nil, errors.New(str)
	}

	return account.Credentials(), nil
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

// FetchQuery fetches a query by name
func (obj *frame) FetchQuery(name string) (queries.Query, error) {
	assignable, err := obj.Fetch(name)
	if err != nil {
		return nil, err
	}

	if !assignable.IsQuery() {
		str := fmt.Sprintf("the assignable (name: %s) was expected to contain a Query", name)
		return nil, errors.New(str)
	}

	return assignable.Query(), nil
}

// HasAssignments returns true if there is assignments, false otherwise
func (obj *frame) HasAssignments() bool {
	return obj.assignments != nil
}

// Assignments fetches the assignments, if any
func (obj *frame) Assignments() Assignments {
	return obj.assignments
}
