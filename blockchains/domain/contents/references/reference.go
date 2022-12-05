package references

type reference struct {
	content    Content
	blockchain Blockchain
}

func createReference(
	content Content,
) Reference {
	return createReferenceInternally(content, nil)
}

func createReferenceWithBlockchain(
	content Content,
	blockchain Blockchain,
) Reference {
	return createReferenceInternally(content, blockchain)
}

func createReferenceInternally(
	content Content,
	blockchain Blockchain,
) Reference {
	out := reference{
		content:    content,
		blockchain: blockchain,
	}

	return &out
}

// Content returns the content
func (obj *reference) Content() Content {
	return obj.content
}

// HasBlockchain returns true if there is a blockchain, false otherwise
func (obj *reference) HasBlockchain() bool {
	return obj.blockchain != nil
}

// Blockchain returns the blockchain, if any
func (obj *reference) Blockchain() Blockchain {
	return obj.blockchain
}
