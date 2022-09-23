package grammars

type grammar struct {
	root     Token
	channels Channels
}

func createGrammar(
	root Token,
) Grammar {
	return createGrammarInternally(root, nil)
}

func createGrammarWithChannels(
	root Token,
	channels Channels,
) Grammar {
	return createGrammarInternally(root, channels)
}

func createGrammarInternally(
	root Token,
	channels Channels,
) Grammar {
	out := grammar{
		root:     root,
		channels: channels,
	}

	return &out
}

// Root returns the root token
func (obj *grammar) Root() Token {
	return obj.root
}

// HasChannels returns true if there is channels, false otherwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *grammar) Channels() Channels {
	return obj.channels
}
