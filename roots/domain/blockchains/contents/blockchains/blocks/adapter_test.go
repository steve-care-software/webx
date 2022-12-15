package blocks

import (
	"bytes"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	block := NewBlockForTests(false)
	adapter := NewAdapter()
	content, err := adapter.ToContent(block)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBlock, err := adapter.ToBlock(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(block.Hash().Bytes(), retBlock.Hash().Bytes()) != 0 {
		t.Errorf("the hash was expected to be '%s', '%s' returned", block.Hash().String(), retBlock.Hash().String())
		return
	}

	if block.Height() != retBlock.Height() {
		t.Errorf("the block was expected to be '%d', '%d' returned", block.Height(), retBlock.Height())
		return
	}

	if bytes.Compare(block.NextScore().Bytes(), retBlock.NextScore().Bytes()) != 0 {
		t.Errorf("the nextScore was expected to be '%s', '%s' returned", block.NextScore().String(), retBlock.NextScore().String())
		return
	}

	if bytes.Compare(block.PendingScore().Bytes(), retBlock.PendingScore().Bytes()) != 0 {
		t.Errorf("the pendingScore was expected to be '%s', '%s' returned", block.PendingScore().String(), retBlock.PendingScore().String())
		return
	}

	if bytes.Compare(block.Transactions().Head().Bytes(), retBlock.Transactions().Head().Bytes()) != 0 {
		t.Errorf("the trx hashTree's head hash was expected to be '%s', '%s' returned", block.Transactions().Head().String(), retBlock.Transactions().Head().String())
		return
	}

	if retBlock.HasPrevious() {
		t.Errorf("the returned block was NOT expected to contain a previous hash")
		return
	}
}

func TestAdapter_withPrevious_Success(t *testing.T) {
	block := NewBlockForTests(true)
	adapter := NewAdapter()
	content, err := adapter.ToContent(block)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBlock, err := adapter.ToBlock(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(block.Hash().Bytes(), retBlock.Hash().Bytes()) != 0 {
		t.Errorf("the hash was expected to be '%s', '%s' returned", block.Hash().String(), retBlock.Hash().String())
		return
	}

	if block.Height() != retBlock.Height() {
		t.Errorf("the block was expected to be '%d', '%d' returned", block.Height(), retBlock.Height())
		return
	}

	if bytes.Compare(block.NextScore().Bytes(), retBlock.NextScore().Bytes()) != 0 {
		t.Errorf("the nextScore was expected to be '%s', '%s' returned", block.NextScore().String(), retBlock.NextScore().String())
		return
	}

	if bytes.Compare(block.PendingScore().Bytes(), retBlock.PendingScore().Bytes()) != 0 {
		t.Errorf("the pendingScore was expected to be '%s', '%s' returned", block.PendingScore().String(), retBlock.PendingScore().String())
		return
	}

	if bytes.Compare(block.Transactions().Head().Bytes(), retBlock.Transactions().Head().Bytes()) != 0 {
		t.Errorf("the trx hashTree's head hash was expected to be '%s', '%s' returned", block.Transactions().Head().String(), retBlock.Transactions().Head().String())
		return
	}

	if bytes.Compare(block.Previous().Bytes(), retBlock.Previous().Bytes()) != 0 {
		t.Errorf("the previous hash was expected to be '%s', '%s' returned", block.Previous().String(), retBlock.Previous().String())
		return
	}
}
