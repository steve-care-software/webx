package trees

import "github.com/steve-care-software/syntax/domain/bytes/grammars"

type element struct {
	grammar   grammars.Element
	content   Content
	amount    uint
	isChannel bool
}

func createElement(
	grammar grammars.Element,
	content Content,
	amount uint,
	isChannel bool,
) Element {
	out := element{
		grammar:   grammar,
		content:   content,
		amount:    amount,
		isChannel: isChannel,
	}

	return &out
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Element {
	return obj.grammar
}

// Content returns the content
func (obj *element) Content() Content {
	return obj.content
}

// Amount returns the amount
func (obj *element) Amount() uint {
	return obj.amount
}

// IsChannel returns true if the element is from a channel, false otherwise
func (obj *element) IsChannel() bool {
	return obj.isChannel
}
