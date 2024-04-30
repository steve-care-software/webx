package failures

const (
	// CouldNotInsertAccountInDatabase represents a could not insert account in database
	CouldNotInsertAccountInDatabase (uint) = iota

	// CouldNotUpdateAccountInDatabase represents a could not update account in database
	CouldNotUpdateAccountInDatabase

	// CouldNotDeleteAccountFromDatabase represents a could not delete account from database
	CouldNotDeleteAccountFromDatabase

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

	// InstanceDoesNotExistsInDatabase represents an instance does not exists in database
	InstanceDoesNotExistsInDatabase

	// InstanceAlreadyExistsInDatabase represents an instance already exists in database
	InstanceAlreadyExistsInDatabase

	// CouldNotFetchConditionFromFrame represents a could not fetch condition from frame
	CouldNotFetchConditionFromFrame

	// CouldNotFetchMessageFromFrame represents a could not fetch message from frame
	CouldNotFetchMessageFromFrame

	// CouldNotFetchAccountFromFrame represents a could not fetch account from frame
	CouldNotFetchAccountFromFrame

	// CouldNotFetchRingFromFrame represents a could not fetch ring from frame
	CouldNotFetchRingFromFrame

	// CouldNotFetchGenerateRingFromFrame represents a could not fetch generate ring from frame
	CouldNotFetchGenerateRingFromFrame

	// CouldNotFetchCipherFromFrame represents a could not fetch cipher from frame
	CouldNotFetchCipherFromFrame

	// CouldNotDecryptCipher represents a could not decrypt cipher
	CouldNotDecryptCipher

	// CouldNotEncryptMessage represents a could not encrypt message
	CouldNotEncryptMessage

	// CouldNotRetrieveAccountFromDatabase represents a could not retrieve account from database
	CouldNotRetrieveAccountFromDatabase

	// CouldNotRetrieveAccountNamesListFromDatabase represents a could not retrieve account names list from database
	CouldNotRetrieveAccountNamesListFromDatabase

	// CouldNotFetchJoinVariableFromFrame represents a could not fetch a join variable from frame
	CouldNotFetchJoinVariableFromFrame

	// CouldNotFetchCompareVariableFromFrame represents a could not fetch a compare variable from frame
	CouldNotFetchCompareVariableFromFrame

	// CouldNotFetchHashVariableFromFrame represents a could not fetch a hash variable from frame
	CouldNotFetchHashVariableFromFrame

	// CouldNotFetchCompileFromFrame represents a could not fetch a compile from frame
	CouldNotFetchCompileFromFrame

	// CouldNotFetchDecompileFromFrame represents a could not fetch a decompile from frame
	CouldNotFetchDecompileFromFrame

	// CouldNotCompileBytesToInstance represents a could not compile bytes to instance
	CouldNotCompileBytesToInstance

	// CouldNotDecompileInstanceToBytes represents a could not fetch a decompile instance to bytes
	CouldNotDecompileInstanceToBytes

	// CouldNotFetchListQueryFromFrame represents a could not fetch list query from frame
	CouldNotFetchListQueryFromFrame

	// CouldNotFetchRetrieveQueryFromFrame represents a could not fetch retrieve query from frame
	CouldNotFetchRetrieveQueryFromFrame

	// CouldNotListInstancesFromDatabase represents a could not list instances from database
	CouldNotListInstancesFromDatabase

	// CouldNotRetrieveInstanceFromDatabase represents a could not retrieve instance from database
	CouldNotRetrieveInstanceFromDatabase

	// CouldNotBeginTransactionInDatabase represents a could not begin transaction in database
	CouldNotBeginTransactionInDatabase

	// CouldNotFetchQueryFromFrame represents a could not fetch query from frame
	CouldNotFetchQueryFromFrame

	// CouldNotCancelInDatabase represents a could not cancel in database service
	CouldNotCancelInDatabase

	// CouldNotCommitInDatabase represents a could not commit in database service
	CouldNotCommitInDatabase

	// CouldNotFetchIndexFromFrame could not fetch index from frame
	CouldNotFetchIndexFromFrame

	// CouldNotRevertToIndexInDatabase could not revert to index in database service
	CouldNotRevertToIndexInDatabase

	// CouldNotRevertInDatabase could not revert in database service
	CouldNotRevertInDatabase

	// CouldNotFetchInstanceFromFrame could not fetch instance from frame
	CouldNotFetchInstanceFromFrame

	// CouldNotInsertInDatabase could not insert in database service
	CouldNotInsertInDatabase

	// CouldNotDeleteFromDatabase could not delete from database service
	CouldNotDeleteFromDatabase

	// CouldNotFetchEncryptionPrivateKeyFromFrame could not fetch an encryption private key from frame
	CouldNotFetchEncryptionPrivateKeyFromFrame

	// CouldNotFetchEncryptionPublicKeyFromFrame could not fetch an encryption public key from frame
	CouldNotFetchEncryptionPublicKeyFromFrame

	// CouldNotFetchSignerPrivateKeyFromFrame could not fetch a signer private key from frame
	CouldNotFetchSignerPrivateKeyFromFrame

	// CouldNotFetchSignerPublicKeyFromFrame could not fetch a signer public key from frame
	CouldNotFetchSignerPublicKeyFromFrame

	// CouldNotSignMessageInFrame could not sign message in frame
	CouldNotSignMessageInFrame

	// CouldNotVoteOnMessageInFrame could not vote on message in frame
	CouldNotVoteOnMessageInFrame

	// CouldNotFetchVoteInFrame could not fetch vote in frame
	CouldNotFetchVoteFromFrame

	// CouldNotFetchHashFromList represents a could not fetch a hash from list
	CouldNotFetchHashFromList
)
