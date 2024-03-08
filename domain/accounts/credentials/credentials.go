package credentials

type credentials struct {
	username string
	password []byte
}

func createCredentials(
	username string,
	password []byte,
) Credentials {
	out := credentials{
		username: username,
		password: password,
	}

	return &out
}

// Username returns the username
func (obj *credentials) Username() string {
	return obj.username
}

// Password returns the password
func (obj *credentials) Password() []byte {
	return obj.password
}
