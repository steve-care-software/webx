package failures

const (
	// CouldNotInsertAccount represents a could not insert account failure
	CouldNotInsertAccount (uint) = iota

	// CouldNotUpdateAccount represents a could not update account
	CouldNotUpdateAccount

	// CouldNotDeleteAccount represents a could not delete account
	CouldNotDeleteAccount

	// CouldNotFetchCredentialsFromFrame represents a could not fetch credentials from frame
	CouldNotFetchCredentialsFromFrame

	// CouldNotFetchUsernameFromFrame represents a could not fetch username from frame
	CouldNotFetchUsernameFromFrame

	// CouldNotFetchPasswordFromFrame represents a could not fetch password from frame
	CouldNotFetchPasswordFromFrame

	// AccountWithSameUsernameAlreadyExists represents an account with same username already exists
	AccountWithSameUsernameAlreadyExists
)
