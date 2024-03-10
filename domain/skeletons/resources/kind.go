package resources

type kind struct {
	native     Native
	reference  []string
	connection string
}

func createKindWithNative(
	native Native,
) Kind {
	return createKindInternally(native, nil, "")
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
	native Native,
	reference []string,
	connection string,
) Kind {
	out := kind{
		native:     native,
		reference:  reference,
		connection: connection,
	}

	return &out
}

// IsNative returns true if there is a native, false otherwise
func (obj *kind) IsNative() bool {
	return obj.native != nil
}

// Native returns the native, if any
func (obj *kind) Native() Native {
	return obj.native
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
