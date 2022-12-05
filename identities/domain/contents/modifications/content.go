package modifications

type content struct {
	name       string
	signature  []byte
	encryption []byte
}

func createContentWithName(
	name string,
) Content {
	return createContentInternally(name, nil, nil)
}

func createContentWithSignaturePK(
	signature []byte,
) Content {
	return createContentInternally("", signature, nil)
}

func createContentWithEncryptionPK(
	encryption []byte,
) Content {
	return createContentInternally("", nil, encryption)
}

func createContentWithNameAndSignaturePK(
	name string,
	signature []byte,
) Content {
	return createContentInternally(name, signature, nil)
}

func createContentWithNameAndEncryptionPK(
	name string,
	encryption []byte,
) Content {
	return createContentInternally(name, nil, encryption)
}

func createContentWithSignaturePKAndEncryptionPK(
	signature []byte,
	encryption []byte,
) Content {
	return createContentInternally("", signature, encryption)
}

func createContentWithNameAndSignaturePKAndEncryptionPK(
	name string,
	signature []byte,
	encryption []byte,
) Content {
	return createContentInternally(name, signature, encryption)
}

func createContentInternally(
	name string,
	signature []byte,
	encryption []byte,
) Content {
	out := content{
		name:       name,
		signature:  signature,
		encryption: encryption,
	}

	return &out
}

// HasName returns true if there is a name, false otherwise
func (obj *content) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *content) Name() string {
	return obj.name
}

// HasSignature returns true if there is a sigPK, false otherwise
func (obj *content) HasSignature() bool {
	return obj.signature != nil
}

// Signature returns the sigPK, if any
func (obj *content) Signature() []byte {
	return obj.signature
}

// HasEncryption returns true if there is an encPK, false otherwise
func (obj *content) HasEncryption() bool {
	return obj.encryption != nil
}

// Encryption returns the encPK, if any
func (obj *content) Encryption() []byte {
	return obj.encryption
}
