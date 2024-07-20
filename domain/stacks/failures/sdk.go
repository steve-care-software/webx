package failures

const (
	// CouldNotFetchPasswordFromFrame represents a could not fetch password from frame
	CouldNotFetchPasswordFromFrame (uint) = iota

	// CouldNotFetchConditionFromFrame represents a could not fetch condition from frame
	CouldNotFetchConditionFromFrame

	// CouldNotFetchMessageFromFrame represents a could not fetch message from frame
	CouldNotFetchMessageFromFrame

	// CouldNotFetchRingFromFrame represents a could not fetch ring from frame
	CouldNotFetchRingFromFrame

	// CouldNotFetchCipherFromFrame represents a could not fetch cipher from frame
	CouldNotFetchCipherFromFrame

	// CouldNotDecryptCipher represents a could not decrypt cipher
	CouldNotDecryptCipher

	// CouldNotEncryptMessage represents a could not encrypt message
	CouldNotEncryptMessage

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

	// CouldNotFetchListFromFrame represents a could not fetch list from frame
	CouldNotFetchListFromFrame

	// CouldNotFetchStringFromFrame represents a could not fetch string from frame
	CouldNotFetchStringFromFrame

	// CouldNotFetchUnsignedIntegerFromFrame represents a could not fetch an unsigned integer from frame
	CouldNotFetchUnsignedIntegerFromFrame

	// CouldNotFetchDeleteFromFrame represents a could not fetch delete from frame
	CouldNotFetchDeleteFromFrame

	// CouldNotFetchBytesFromFrame represents a could not fetch bytes from frame
	CouldNotFetchBytesFromFrame

	// CouldNotFetchBoolFromFrame represents a could not fetch bool from frame
	CouldNotFetchBoolFromFrame

	// CouldNotFetchCommitFromFrame represents a could not fetch commit from frame
	CouldNotFetchCommitFromFrame

	// CouldNotFetchEncryptionPrivateKeyFromFrame could not fetch an encryption private key from frame
	CouldNotFetchEncryptionPrivateKeyFromFrame

	// CouldNotFetchEncryptionPublicKeyFromFrame could not fetch an encryption public key from frame
	CouldNotFetchEncryptionPublicKeyFromFrame

	// CouldNotFetchSignerPrivateKeyFromFrame could not fetch a signer private key from frame
	CouldNotFetchSignerPrivateKeyFromFrame

	// CouldNotFetchSignerPublicKeyFromFrame could not fetch a signer public key from frame
	CouldNotFetchSignerPublicKeyFromFrame

	// CouldNotFetchSignatureFromFrame could not fetch a signature from frame
	CouldNotFetchSignatureFromFrame

	// CouldNotVoteOnMessageInFrame could not vote on message in frame
	CouldNotVoteOnMessageInFrame

	// CouldNotFetchVoteInFrame could not fetch vote in frame
	CouldNotFetchVoteFromFrame

	// CouldNotFetchHashFromList represents a could not fetch a hash from list
	CouldNotFetchHashFromList

	// CouldNotFetchStringFromList represents a could not fetch a string from list
	CouldNotFetchStringFromList

	// CouldNotFetchModificationFromList represents a could not fetch a modification from list
	CouldNotFetchModificationFromList

	// CouldNotFetchActionFromList represents a could not fetch an action from list
	CouldNotFetchActionFromList

	// CouldNotRetrieveListFromRepository represents a could not retrieve list from repository
	CouldNotRetrieveListFromRepository

	// CouldNotRetrieveEmptyListFromRepository represents a could not retrieve an empty list from repository
	CouldNotRetrieveEmptyListFromRepository

	// CouldNotRetrieveFromRepository represents a could not retrieve from repository
	CouldNotRetrieveFromRepository

	// CouldNotExecuteExistsFromRepository represents a could not execute exists from repository
	CouldNotExecuteExistsFromRepository

	// CouldNotFetchElementFromList represents a could not fetch element from list
	CouldNotFetchElementFromList

	// CouldNotFetchFromFrame represents a could not fetch from frame
	CouldNotFetchFromFrame

	// CouldNotDeleteDatabaseFromService represents a could not delete database from service
	CouldNotDeleteDatabaseFromService

	// CouldNotSaveDatabaseFromService represents a could not dave database from service
	CouldNotSaveDatabaseFromService

	// CouldNotFetchDatabaseFromFrame represents a could not fetch database from frame
	CouldNotFetchDatabaseFromFrame

	// CouldNotFetchExecutableeFromFrame represents a could not fetch executable from frame
	CouldNotFetchExecutableeFromFrame

	// CouldNotExecuteListFromExecutable represents a could not execute list from executable
	CouldNotExecuteListFromExecutable

	// CouldNotExecuteBeginFromExecutable represents a could not execute begin from executable
	CouldNotExecuteBeginFromExecutable

	// CouldNotExecuteAmountFromExecutable represents a could not execute amount from executable
	CouldNotExecuteAmountFromExecutable

	// CouldNotExecuteHeadFromExecutable represents a could not execute head from executable
	CouldNotExecuteHeadFromExecutable

	// CouldNotExecuteExecuteFromExecutable represents a could not execute execute from executable
	CouldNotExecuteExecuteFromExecutable

	// CouldNotExecuteWithPathFromExecutable represents a could not execute executeWithPath from executable
	CouldNotExecuteWithPathFromExecutable

	// CouldNotExecuteExecuteLayerFromExecutable represents a could not execute executeLayer from executable
	CouldNotExecuteExecuteLayerFromExecutable

	// CouldNotExecuteExecuteLayerWithPathFromExecutable represents a could not execute executeLayerWithPath from executable
	CouldNotExecuteExecuteLayerWithPathFromExecutable

	// CouldNotExecuteInithFromExecutable represents a could not execute init from executable
	CouldNotExecuteInitFromExecutable

	// CouldNotExecuteRetrieveAtFromExecutable represents a could not execute retrieveAt from executable
	CouldNotExecuteRetrieveAtFromExecutable

	// CouldNotExecuteRetrieveAllFromExecutable represents a could not execute retrieveAll from executable
	CouldNotExecuteRetrieveAllFromExecutable
)
