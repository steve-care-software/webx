package resources

type kind struct {
	pNative    *uint8
	reference  []string
	connection string
}

func createKindWithNative(
	pNative *uint8,
) Kind {
	return createKindInternally(pNative, nil, "")
}

func createKindWithReference(
	reference []string,
) Kind {
	return createKindInternally(nil, reference, "")
}

func createKindWithConnection(
	connection string,
) Kind {
	return createKindInternally(nil, nil, connection)
}

func createKindInternally(
	pNative *uint8,
	reference []string,
	connection string,
) Kind {
	out := kind{
		pNative:    pNative,
		reference:  reference,
		connection: connection,
	}

	return &out
}

// IsNative returns true if there is a native, false otherwise
func (obj *kind) IsNative() bool {
	return obj.pNative != nil
}

// Native returns the native, if any
func (obj *kind) Native() *uint8 {
	return obj.pNative
}

// IsReference returns true if there is a reference, false otherwise
func (obj *kind) IsReference() bool {
	return obj.reference != nil
}

// Reference returns the reference, if any
func (obj *kind) Reference() []string {
	return obj.reference
}

// IsConnection returns true if there is a connection, false otherwise
func (obj *kind) IsConnection() bool {
	return obj.connection != ""
}

// Connection returns the connection, if any
func (obj *kind) Connection() string {
	return obj.connection
}
