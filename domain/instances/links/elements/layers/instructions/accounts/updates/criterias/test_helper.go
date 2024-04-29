package criterias

// NewCriteriaForTests creates a new criteria for tests
func NewCriteriaForTests(changeSigner bool, changeEncryptor bool) Criteria {
	return newCriteriaForTests(changeSigner, changeEncryptor, "", "")
}

// NewCriteriaWithUsernameForTests creates a new criteria with username for tests
func NewCriteriaWithUsernameForTests(changeSigner bool, changeEncryptor bool, username string) Criteria {
	return newCriteriaForTests(changeSigner, changeEncryptor, username, "")
}

// NewCriteriaWithPasswordForTests creates a new criteria with password for tests
func NewCriteriaWithPasswordForTests(changeSigner bool, changeEncryptor bool, password string) Criteria {
	return newCriteriaForTests(changeSigner, changeEncryptor, "", password)
}

// NewCriteriaWithUsernameAndPasswordForTests creates a new criteria with username and password for tests
func NewCriteriaWithUsernameAndPasswordForTests(changeSigner bool, changeEncryptor bool, username string, password string) Criteria {
	return newCriteriaForTests(changeSigner, changeEncryptor, username, password)
}

func newCriteriaForTests(changeSigner bool, changeEncryptor bool, username string, password string) Criteria {
	builder := NewBuilder().Create()
	if changeSigner {
		builder.ChangeSigner()
	}

	if changeEncryptor {
		builder.ChangeEncryptor()
	}

	if username != "" {
		builder.WithUsername(username)
	}

	if password != "" {
		builder.WithPassword(password)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
