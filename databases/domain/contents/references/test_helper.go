package references

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/steve-care-software/webx/databases/domain/contents/peers"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// NewReferenceForTests creates a new reference with peers for tests
func NewReferenceForTests(maxPeerAmount uint) Reference {
	contentKeys, err := NewContentKeysBuilder().Create().WithList([]ContentKey{
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
	}).Now()
	if err != nil {
		panic(err)
	}

	commits := NewCommitsForTests(32)
	builder := NewBuilder().Create().WithContentKeys(contentKeys).WithCommits(commits)

	peersList := peers.NewPeersForTests(maxPeerAmount)
	if len(peersList) > 0 {
		builder.WithPeers(peersList)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitsForTests creates a new commits for tests
func NewCommitsForTests(amount uint) Commits {
	list := []Commit{}
	for i := 0; i < int(amount); i++ {
		list = append(list, NewCommitForTests())
	}

	ins, err := NewCommitsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitForTests creates a new commit for tests
func NewCommitForTests() Commit {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some bytes: %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	pointer := NewPointerForTests()
	createdOn := time.Now().UTC()
	ins, err := NewCommitBuilder().Create().WithHash(*pHash).WithPointer(pointer).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests() Pointer {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	from := uint(r1.Int() + 1)

	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)

	length := uint(r1.Int() + 1)
	pointer, err := NewPointerBuilder().Create().From(from).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return pointer
}

// NewContentKeyForTests creates a new content key for tests
func NewContentKeyForTests() ContentKey {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some data %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	pCommitHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some commit data %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	from := uint(r1.Intn(233456))
	length := uint(r1.Intn(22323)) + 1
	pointer, err := NewPointerBuilder().Create().From(from).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	ins, err := NewContentKeyBuilder().Create().
		WithHash(*pHash).
		WithKind(43).
		WithContent(pointer).
		WithCommit(*pCommitHash).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
