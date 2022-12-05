package channels

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	channel := NewChannelForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(channel)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retChannel, err := adapter.ToChannel(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(channel, retChannel) {
		t.Errorf("the returned channel is invalid")
		return
	}
}

func TestAdapter_withPrevious_Success(t *testing.T) {
	channel := NewChannelWithPreviousForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(channel)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retChannel, err := adapter.ToChannel(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(channel, retChannel) {
		t.Errorf("the returned channel is invalid")
		return
	}
}

func TestAdapter_withNext_Success(t *testing.T) {
	channel := NewChannelWithNextForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(channel)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retChannel, err := adapter.ToChannel(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(channel, retChannel) {
		t.Errorf("the returned channel is invalid")
		return
	}
}

func TestAdapter_withPrevious_withNext_Success(t *testing.T) {
	channel := NewChannelWithPreviousAndNextForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(channel)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retChannel, err := adapter.ToChannel(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(channel, retChannel) {
		t.Errorf("the returned channel is invalid")
		return
	}
}
