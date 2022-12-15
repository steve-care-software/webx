package hashtrees

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// we must also split data, create a tree, create a compact tree, and pass the shuffled data to it, to get it back in order
// when passing an invalid amount of blocks to the CreateHashTree, returns an error (1, for example.)
func createTreeAndTest(t *testing.T, text string, delimiter string, height uint) {
	adapter := NewAdapter()
	shuf := func(v [][]byte) {
		f := reflect.Swapper(v)
		n := len(v)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < n; i++ {
			f(r.Intn(n), r.Intn(n))
		}
	}

	splittedData := bytes.Split([]byte(text), []byte(delimiter))
	splittedDataLength := len(splittedData)
	splittedDataLengthPowerOfTwo := uint(math.Pow(2, math.Ceil(math.Log(float64(splittedDataLength))/math.Log(2))))
	tree, err := NewBuilder().Create().WithBlocks(splittedData).Now()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", err.Error())
		return
	}

	if tree == nil {
		t.Errorf("the returned instance was expected to be an instance, nil returned")
		return
	}

	secondTree, err := NewBuilder().Create().WithBlocks(splittedData).Now()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", err.Error())
		return
	}

	if tree.Head().String() != secondTree.Head().String() {
		t.Errorf("the tree hashes changed even if they were build with the same data: First: %s, Second: %s", tree.Head().String(), secondTree.Head().String())
		return
	}

	treeHeight := tree.Height()
	if treeHeight != height {
		t.Errorf("the binary tree's height should be %d because it contains %d data blocks, %d given", height, len(splittedData), treeHeight)
		return
	}

	pTreeLength, err := adapter.ToLength(tree)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", err.Error())
		return
	}

	if *pTreeLength != splittedDataLengthPowerOfTwo {
		t.Errorf("the HashTree should have a length of %d, %d given", splittedDataLengthPowerOfTwo, *pTreeLength)
		return
	}

	compact, err := adapter.ToCompact(tree)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", err.Error())
		return
	}

	compactLength := compact.Length()
	if splittedDataLengthPowerOfTwo != compactLength {
		t.Errorf("the CompactHashTree should have a length of %d, %d given", splittedDataLengthPowerOfTwo, compactLength)
		return
	}

	if !tree.Head().Compare(compact.Head()) {
		t.Errorf("the HashTree root hash: %x is not the same as the CompactHashTree root hash: %x", tree.Head().Bytes(), compact.Head().Bytes())
		return
	}

	shuffledData := make([][]byte, len(splittedData))
	copy(shuffledData, splittedData)
	shuf(shuffledData)

	reOrderedSplittedData, err := adapter.ToOrder(tree, shuffledData)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(splittedData, reOrderedSplittedData) {
		t.Errorf("the re-ordered data is invalid")
		return
	}

	treeData, err := adapter.ToContent(tree)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTree, err := adapter.ToHashTree(treeData)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(tree, retTree) {
		t.Errorf("the returned tree is invalid")
		return
	}
}

func TestHashTree_Success(t *testing.T) {
	createTreeAndTest(t, "this|is", "|", 2)                                                                                                                       //2 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf", "|", 4)                                                                               //8 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|another", "|", 5)                                                                       //9 blocks, rounded up to 16
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|another|lol", "|", 5)                                                                   //10 blocks, rounded up to 16
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|asfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasfd|sdfasd", "|", 5)          //16 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|asfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasfd|sdfasd|dafgsagf", "|", 6) //17 blocks, rounded up to 32
}

func TestHashTree_withOneBlock_returnsError(t *testing.T) {
	adapter := NewAdapter()

	//variables:
	text := "this"
	delimiter := "|"

	splittedData := bytes.Split([]byte(text), []byte(delimiter))
	tree, err := NewBuilder().Create().WithBlocks(splittedData).Now()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	orderedData, err := adapter.ToOrder(tree, splittedData)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(orderedData, splittedData) {
		t.Errorf("the ordered data was invalid")
		return
	}

	if tree.Height() != 2 {
		t.Errorf("the height of the tree was edxpected to be 2, %d returned", tree.Height())
		return
	}

	pTreeLength, err := adapter.ToLength(tree)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if *pTreeLength != 2 {
		t.Errorf("the length of the tree was edxpected to be 2, %d returned", *pTreeLength)
		return
	}
}

func TestHashTree_convertToJSON_backAndForth_Success(t *testing.T) {

	adapter := NewAdapter()
	leavesAdapter := NewLeavesAdapter()
	compactAdapter := NewCompactAdapter()

	//variables:
	r := rand.New(rand.NewSource(99))
	blks := [][]byte{
		[]byte("this"),
		[]byte("is"),
		[]byte("some"),
		[]byte("blocks"),
		[]byte(fmt.Sprintf("some rand number to make it unique: %d", r.Int())),
	}

	//execute:
	hashTree, err := NewBuilder().Create().WithBlocks(blks).Now()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if hashTree == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

	compact, err := adapter.ToCompact(hashTree)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
	}

	retHashTree, err := leavesAdapter.ToHashTree(compact.Leaves())
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
	}

	if !reflect.DeepEqual(hashTree, retHashTree) {
		t.Errorf("the returned hashTree is invalid")
	}

	compactData, err := compactAdapter.ToContent(compact)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
	}

	retCompact, err := compactAdapter.ToCompact(compactData)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
	}

	if !reflect.DeepEqual(compact, retCompact) {
		t.Errorf("the returned compact is invalid")
	}
}
