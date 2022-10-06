package grammars

type instance struct {
	token      Token
	everything Everything
}

func createInstanceWithToken(
	token Token,
) Instance {
	return createInstanceInternally(token, nil)
}

func createInstanceWithEverything(
	everything Everything,
) Instance {
	return createInstanceInternally(nil, everything)
}

func createInstanceInternally(
	token Token,
	everything Everything,
) Instance {
	out := instance{
		token:      token,
		everything: everything,
	}

	return &out
}

// Name returns the name
func (obj *instance) Name() string {
	if obj.IsToken() {
		return obj.Token().Name()
	}

	return obj.Everything().Name()
}

// IsToken returns true if there is a token, false otherwise
func (obj *instance) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *instance) Token() Token {
	return obj.token
}

// IsEverything returns true if there is an everything, false otherwise
func (obj *instance) IsEverything() bool {
	return obj.everything != nil
}

// Everything returns the everything, if any
func (obj *instance) Everything() Everything {
	return obj.everything
}
