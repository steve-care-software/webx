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

	// CouldNotFetchContextFromFrame represents a could not fetch context from frame
	CouldNotFetchContextFromFrame

	// CouldNotFetchPathFromFrame represents a could not fetch path from frame
	CouldNotFetchPathFromFrame

	// CouldNotFetchIdentifierFromFrame represents a could not fetch identifier from frame
	CouldNotFetchIdentifierFromFrame

	// InstanceDoesNotExists represents an insatnce does not exists
	InstanceDoesNotExists

	// InstanceAlreadyExists represents an insatnce already exists
	InstanceAlreadyExists

	// CouldNotFetchConditionFromFrame represents a could not fetch condition from frame
	CouldNotFetchConditionFromFrame
)
