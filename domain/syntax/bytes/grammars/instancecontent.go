package grammars

type instanceContent struct {
	token      Token
	everything Everything
}

func createInstanceContentWithToken(
	token Token,
) InstanceContent {
	return createInstanceContentInternally(token, nil)
}

func createInstanceContentWithEverything(
	everything Everything,
) InstanceContent {
	return createInstanceContentInternally(nil, everything)
}

func createInstanceContentInternally(
	token Token,
	everything Everything,
) InstanceContent {
	out := instanceContent{
		token:      token,
		everything: everything,
	}

	return &out
}

// IsToken returns true if there is a token, false otherwise
func (obj *instanceContent) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *instanceContent) Token() Token {
	return obj.token
}

// IsEverything returns true if there is an everything, false otherwise
func (obj *instanceContent) IsEverything() bool {
	return obj.everything != nil
}

// Everything returns the everything, if any
func (obj *instanceContent) Everything() Everything {
	return obj.everything
}
