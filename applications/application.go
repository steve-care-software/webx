package logics

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/datastencil/domain/receipts"
	"github.com/steve-care-software/datastencil/domain/receipts/commands"
	"github.com/steve-care-software/datastencil/domain/receipts/commands/results"
	"github.com/steve-care-software/datastencil/domain/stacks"
	accounts_applications "github.com/steve-care-software/identity/applications"
	"github.com/steve-care-software/identity/domain/accounts"
	account_encryptors "github.com/steve-care-software/identity/domain/accounts/encryptors"
	"github.com/steve-care-software/identity/domain/accounts/signers"
	account_signers "github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/identity/domain/credentials"
	"github.com/steve-care-software/identity/domain/hash"
)

type application struct {
	accountApplication accounts_applications.Application
	hashAdapter        hash.Adapter
	signerFactory      signers.Factory
	signerVoteAdapter  signers.VoteAdapter
	receiptBuilder     receipts.ReceiptBuilder
	commandsBuilder    commands.Builder
	commandBuilder     commands.CommandBuilder
	resultBuilder      results.Builder
	successBuilder     results.SuccessBuilder
	failureBuilder     results.FailureBuilder
	stackFactory       stacks.Factory
	stackBuilder       stacks.Builder
	framesBuilder      stacks.FramesBuilder
	frameBuilder       stacks.FrameBuilder
	assignmentsBuilder stacks.AssignmentsBuilder
	assignmentBuilder  stacks.AssignmentBuilder
	assignableBuilder  stacks.AssignableBuilder
	credentials        credentials.Credentials
}

func createApplication(
	accountApplication accounts_applications.Application,
	hashAdapter hash.Adapter,
	signerFactory signers.Factory,
	signerVoteAdapter signers.VoteAdapter,
	receiptBuilder receipts.ReceiptBuilder,
	commandsBuilder commands.Builder,
	commandBuilder commands.CommandBuilder,
	resultBuilder results.Builder,
	successBuilder results.SuccessBuilder,
	failureBuilder results.FailureBuilder,
	stackFactory stacks.Factory,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	assignmentBuilder stacks.AssignmentBuilder,
	assignableBuilder stacks.AssignableBuilder,
	credentials credentials.Credentials,
) Application {
	out := application{
		accountApplication: accountApplication,
		hashAdapter:        hashAdapter,
		signerFactory:      signerFactory,
		signerVoteAdapter:  signerVoteAdapter,
		receiptBuilder:     receiptBuilder,
		commandsBuilder:    commandsBuilder,
		commandBuilder:     commandBuilder,
		resultBuilder:      resultBuilder,
		successBuilder:     successBuilder,
		failureBuilder:     failureBuilder,
		stackFactory:       stackFactory,
		stackBuilder:       stackBuilder,
		framesBuilder:      framesBuilder,
		frameBuilder:       frameBuilder,
		assignmentsBuilder: assignmentsBuilder,
		assignmentBuilder:  assignmentBuilder,
		assignableBuilder:  assignableBuilder,
		credentials:        credentials,
	}

	return &out
}

// Execute executes the logic application
func (app *application) Execute(input []byte, layer layers.Layer, library libraries.Library, context receipts.Receipt) (receipts.Receipt, error) {
	// authenticate:
	currentAccount, err := app.accountApplication.Retrieve(app.credentials)
	if err != nil {
		return nil, err
	}

	// execute:
	return app.authenticatedExecute(
		currentAccount,
		input, layer,
		library,
		context,
	)
}

func (app *application) authenticatedExecute(
	currentAccount accounts.Account,
	input []byte,
	layer layers.Layer,
	library libraries.Library,
	context receipts.Receipt,
) (receipts.Receipt, error) {
	// execute the layer:
	retReceipt, err := app.executeLayer(currentAccount, input, layer, library, context)
	if err != nil {
		return nil, err
	}

	// retrieve the link related to our executed layers, if any, from our library:
	if library.HasLinks() {
		// build the list of executed layers:
		layerHashes := []hash.Hash{}
		commandsList := retReceipt.Commands().List()
		for _, oneCommand := range commandsList {
			layerHashes = append(layerHashes, oneCommand.Layer().Hash())
		}

		link, err := library.Links().FetchByExecutedLayers(layerHashes)
		if err != nil {
			// no link to execute:
			return retReceipt, nil
		}

		// execute the link:
		return app.executeLink(
			currentAccount,
			link,
			library,
			retReceipt,
		)
	}

	// execute the link:
	return retReceipt, nil
}

func (app *application) executeLink(
	currentAccount accounts.Account,
	link links.Link,
	library libraries.Library,
	previousReceipt receipts.Receipt,
) (receipts.Receipt, error) {
	previousResult := previousReceipt.Commands().Last().Result()
	if previousResult.IsFailure() {
		hash := link.Hash().String()
		str := fmt.Sprintf("the link (hash: %s) cannot execute because the previous result failed", hash)
		return nil, errors.New(str)
	}

	previousSuccess := previousResult.Success()
	currentContext := previousReceipt
	elementsList := link.Elements().List()
	for _, oneElement := range elementsList {
		retReceipt, retSuccess, err := app.executeLinkElement(
			currentAccount,
			oneElement,
			library,
			previousSuccess,
			currentContext,
		)

		if err != nil {
			return nil, err
		}

		currentContext = retReceipt
		previousSuccess = retSuccess
	}

	return currentContext, nil
}

func (app *application) executeLinkElement(
	currentAccount accounts.Account,
	element links.Element,
	library libraries.Library,
	previousSuccess results.Success,
	context receipts.Receipt,
) (receipts.Receipt, results.Success, error) {
	// execute the layer:
	layerHash := element.Layer()
	layer, err := library.Layers().Fetch(layerHash)
	if err != nil {
		return nil, nil, err
	}

	input := previousSuccess.Bytes()
	retReceipt, err := app.authenticatedExecute(currentAccount, input, layer, library, context)
	if err != nil {
		return nil, nil, err
	}

	retSuccess := app.fetchSuccess(context, retReceipt)
	if !element.HasCondition() {
		return retReceipt, retSuccess, nil
	}

	result := retReceipt.Commands().Last().Result()
	if result.IsSuccess() {
		return retReceipt, retSuccess, nil
	}

	condition := element.Condition()
	failure := result.Failure()
	if app.matchLinkCondition(condition, failure) {
		return retReceipt, previousSuccess, nil
	}

	str := fmt.Sprintf("the layer (hash: %s) did not execute successfully and the link condition did not match", layerHash)
	return nil, nil, errors.New(str)
}

func (app *application) fetchSuccess(
	previous receipts.Receipt,
	current receipts.Receipt,
) results.Success {
	result := current.Commands().Last().Result()
	if result.IsSuccess() {
		return result.Success()
	}

	return previous.Commands().Last().Result().Success()
}

func (app *application) matchLinkCondition(
	condition links.Condition,
	failure results.Failure,
) bool {
	resource := condition.Resource()
	isMatch := app.matchLinkConditionResource(resource, failure)
	if isMatch {
		return true
	}

	if !condition.HasNext() {
		return false
	}

	next := condition.Next()
	if next.IsResource() {
		resource := next.Resource()
		return app.matchLinkConditionResource(resource, failure)
	}

	nextCondition := next.Condition()
	return app.matchLinkCondition(nextCondition, failure)
}

func (app *application) matchLinkConditionResource(
	resource links.ConditionResource,
	failure results.Failure,
) bool {
	actualCode := failure.Code()
	actualIsRaisedInLayer := failure.IsRaisedInLayer()
	expectedCode := resource.Code()
	expectedIsRaisedInLayer := resource.IsRaisedInLayer()
	return actualCode == expectedCode && actualIsRaisedInLayer == expectedIsRaisedInLayer
}

func (app *application) executeLayer(
	currentAccount accounts.Account,
	input []byte,
	layer layers.Layer,
	library libraries.Library,
	context receipts.Receipt,
) (receipts.Receipt, error) {
	assignable, err := app.assignableBuilder.Create().WithBytes(input).Now()
	if err != nil {
		return nil, err
	}

	variable := layer.Input()
	assignment, err := app.assignmentBuilder.Create().WithName(variable).WithAssignable(assignable).Now()
	if err != nil {
		return nil, err
	}

	assignments, err := app.assignmentsBuilder.Create().WithList([]stacks.Assignment{
		assignment,
	}).Now()

	if err != nil {
		return nil, err
	}

	frame, err := app.frameBuilder.Create().WithAssignments(assignments).Now()
	if err != nil {
		return nil, err
	}

	frames, err := app.framesBuilder.Create().WithList([]stacks.Frame{
		frame,
	}).Now()

	if err != nil {
		return nil, err
	}

	stack, err := app.stackBuilder.Create().WithFrames(frames).Now()
	if err != nil {
		return nil, err
	}

	instructions := layer.Instructions()
	retStack, retFailure, retReceipt, err := app.executeInstructions(
		currentAccount,
		library,
		context,
		layer,
		stack,
		instructions,
	)

	if err != nil {
		return nil, err
	}

	signer := currentAccount.Signer()
	return app.generateReceipt(
		signer,
		input,
		layer,
		retStack,
		retFailure,
		retReceipt,
		nil,
	)
}

func (app *application) generateReceipt(
	currentSigner signers.Signer,
	input []byte,
	layer layers.Layer,
	stack stacks.Stack,
	failure results.Failure,
	receipt receipts.Receipt,
	parent commands.Link,
) (receipts.Receipt, error) {
	commandBuilder := app.commandBuilder.Create().WithInput(input).WithLayer(layer)
	if failure != nil {
		retResult, err := app.resultBuilder.Create().
			WithFailure(failure).
			Now()

		if err != nil {
			return nil, err
		}

		commandBuilder.WithResult(retResult)
	}

	if failure == nil {
		output := layer.Output()
		variable := output.Variable()
		outputBytes, err := stack.Head().FetchBytes(variable)
		if err != nil {
			return nil, err
		}

		if output.HasExecute() {

		}

		kind := output.Kind()
		success, err := app.successBuilder.Create().WithBytes(outputBytes).WithKind(kind).Now()
		if err != nil {
			return nil, err
		}

		retResult, err := app.resultBuilder.Create().
			WithSuccess(success).
			Now()

		if err != nil {
			return nil, err
		}

		commandBuilder.WithResult(retResult)
	}

	if parent != nil {
		commandBuilder.WithParent(parent)
	}

	newCommand, err := commandBuilder.Now()
	if err != nil {
		return nil, err
	}

	commandsList := receipt.Commands().List()
	commandsList = append(commandsList, newCommand)
	commands, err := app.commandsBuilder.Create().WithList(commandsList).Now()
	if err != nil {
		return nil, err
	}

	msg := commands.Hash().Bytes()
	signature, err := currentSigner.Sign(msg)
	if err != nil {
		return nil, err
	}

	return app.receiptBuilder.Create().
		WithCommands(commands).
		WithSignature(signature).
		Now()
}

func (app *application) executeInstructions(
	currentAccount accounts.Account,
	library libraries.Library,
	receipts receipts.Receipt,
	currentLayer layers.Layer,
	stack stacks.Stack,
	instructions layers.Instructions,
) (stacks.Stack, results.Failure, receipts.Receipt, error) {
	var currentFailure results.Failure
	currentStack := stack
	currentContext := receipts
	list := instructions.List()
	for _, oneInstruction := range list {
		stop, retStack, retFailure, retUpdatedReceipts, err := app.executeInstruction(
			currentAccount,
			library,
			currentContext,
			currentLayer,
			currentStack,
			oneInstruction,
		)

		if err != nil {
			return nil, nil, nil, err
		}

		if stop {
			break
		}

		currentStack = retStack
		currentContext = retUpdatedReceipts

		if retFailure != nil {
			currentFailure = retFailure
			break
		}
	}

	return currentStack, currentFailure, currentContext, nil
}

func (app *application) executeInstruction(
	currentAccount accounts.Account,
	library libraries.Library,
	currentContext receipts.Receipt,
	currentLayer layers.Layer,
	stack stacks.Stack,
	instruction layers.Instruction,
) (bool, stacks.Stack, results.Failure, receipts.Receipt, error) {
	headFrame := stack.Head()
	if instruction.IsStop() {
		return true, stack, nil, currentContext, nil
	}

	if instruction.IsRaiseError() {
		code := instruction.RaiseError()
		failure, err := app.executeRaiseError(code)
		if err != nil {
			return false, nil, nil, nil, err
		}

		return false, nil, failure, nil, nil
	}

	if instruction.IsCondition() {
		condition := instruction.Condition()
		retStack, retFailure, retReceipts, err := app.executeCondition(
			currentAccount,
			library,
			currentContext,
			currentLayer,
			stack,
			condition,
		)

		if err != nil {
			return false, nil, nil, nil, err
		}

		return false, retStack, retFailure, retReceipts, nil
	}

	assignment := instruction.Assignment()
	retFrame, retFailure, retReceipts, err := app.executeAssignment(
		currentAccount,
		library,
		currentContext,
		currentLayer,
		headFrame,
		assignment,
	)

	if err != nil {
		return false, nil, nil, nil, err
	}

	framesList := stack.Frames().List()
	framesList = append(framesList, retFrame)
	updatedFrames, err := app.framesBuilder.Create().
		WithList(framesList).
		Now()

	if err != nil {
		return false, nil, nil, nil, err
	}

	updatedStack, err := app.stackBuilder.Create().
		WithFrames(updatedFrames).
		Now()

	if err != nil {
		return false, nil, nil, nil, err
	}

	return false, updatedStack, retFailure, retReceipts, nil
}

func (app *application) executeRaiseError(code uint) (results.Failure, error) {
	failure, err := app.failureBuilder.Create().
		WithCode(code).
		IsRaisedInLayer().
		Now()

	if err != nil {
		return nil, err
	}

	return failure, nil
}

func (app *application) executeCondition(
	currentAccount accounts.Account,
	library libraries.Library,
	currentContext receipts.Receipt,
	currentLayer layers.Layer,
	stack stacks.Stack,
	condition layers.Condition,
) (stacks.Stack, results.Failure, receipts.Receipt, error) {
	variable := condition.Variable()
	boolValue, err := stack.Head().FetchBool(variable)
	if err != nil {

	}

	if boolValue {
		instructions := condition.Instructions()
		return app.executeInstructions(
			currentAccount,
			library,
			currentContext,
			currentLayer,
			stack,
			instructions,
		)
	}

	return stack, nil, currentContext, nil
}

func (app *application) executeAssignment(
	currentAccount accounts.Account,
	library libraries.Library,
	currentContext receipts.Receipt,
	currentLayer layers.Layer,
	frame stacks.Frame,
	assignment layers.Assignment,
) (stacks.Frame, results.Failure, receipts.Receipt, error) {
	assignable := assignment.Assignable()
	retAssignable, failure, receipts, err := app.executeAssignable(
		currentAccount,
		library,
		currentContext,
		currentLayer,
		frame,
		assignable,
	)

	if err != nil {
		return nil, nil, nil, err
	}

	name := assignment.Name()
	currentAssignmentsList := []stacks.Assignment{}
	if frame.HasAssignments() {
		currentAssignmentsList = frame.Assignments().List()
	}

	newAssignment, err := app.assignmentBuilder.Create().
		WithName(name).
		WithAssignable(retAssignable).
		Now()

	if err != nil {
		return nil, nil, nil, err
	}

	currentAssignmentsList = append(currentAssignmentsList, newAssignment)
	assignments, err := app.assignmentsBuilder.Create().
		WithList(currentAssignmentsList).
		Now()

	if err != nil {
		return nil, nil, nil, err
	}

	retFrame, err := app.frameBuilder.Create().
		WithAssignments(assignments).
		Now()

	if err != nil {
		return nil, nil, nil, err
	}

	return retFrame, failure, receipts, nil
}

func (app *application) executeAssignable(
	currentAccount accounts.Account,
	library libraries.Library,
	currentContext receipts.Receipt,
	currentLayer layers.Layer,
	frame stacks.Frame,
	assignable layers.Assignable,
) (stacks.Assignable, results.Failure, receipts.Receipt, error) {
	if assignable.IsBytes() {
		bytesIns := assignable.Bytes()
		retAssignable, err := app.executeBytes(
			frame,
			bytesIns,
		)

		if err != nil {
			return nil, nil, nil, err
		}

		return retAssignable, nil, currentContext, nil
	}

	if assignable.IsIdentity() {
		identity := assignable.Identity()
		retAssignable, err := app.executeIdentity(
			currentAccount,
			frame,
			identity,
		)

		if err != nil {
			return nil, nil, nil, err
		}

		return retAssignable, nil, currentContext, nil
	}

	execution := assignable.Execution()
	return app.executeExecution(
		currentAccount,
		library,
		currentContext,
		currentLayer,
		frame,
		execution,
	)
}

func (app *application) executeExecution(
	currentAccount accounts.Account,
	library libraries.Library,
	currentContext receipts.Receipt,
	currentLayer layers.Layer,
	frame stacks.Frame,
	execution layers.Execution,
) (stacks.Assignable, results.Failure, receipts.Receipt, error) {
	inputVariable := execution.Input()
	input, err := frame.FetchBytes(inputVariable)
	if err != nil {
		return nil, nil, nil, err
	}

	layerToExecute := currentLayer
	if execution.HasLayer() {
		layerHashVariable := execution.Layer()
		layerHash, err := frame.FetchHash(layerHashVariable)
		if err != nil {
			return nil, nil, nil, err
		}

		layer, err := library.Layers().Fetch(layerHash)
		if err != nil {
			return nil, nil, nil, err
		}

		layerToExecute = layer
	}

	retReceipt, err := app.authenticatedExecute(
		currentAccount,
		input,
		layerToExecute,
		library,
		currentContext,
	)

	if err != nil {
		return nil, nil, nil, err
	}

	result := retReceipt.Commands().Last().Result()
	if result.IsSuccess() {
		data := result.Success().Bytes()
		retAssignable, err := app.assignableBuilder.Create().
			WithBytes(data).
			Now()

		if err != nil {
			return nil, nil, nil, err
		}

		return retAssignable, nil, retReceipt, nil
	}

	failure := result.Failure()
	return nil, failure, retReceipt, nil
}

func (app *application) executeIdentity(
	currentAccount accounts.Account,
	frame stacks.Frame,
	identity layers.Identity,
) (stacks.Assignable, error) {
	if identity.IsSigner() {
		signer := identity.Signer()
		accountSigner := currentAccount.Signer()
		return app.executeSigner(
			accountSigner,
			frame,
			signer,
		)
	}

	encryptor := identity.Encryptor()
	accountEncryptor := currentAccount.Encryptor()
	return app.executeEncryptor(
		accountEncryptor,
		frame,
		encryptor,
	)
}

func (app *application) executeSigner(
	accountSigner account_signers.Signer,
	frame stacks.Frame,
	signer layers.Signer,
) (stacks.Assignable, error) {
	if signer.IsBytes() {
		variable := signer.Bytes()
		return app.executeSignerBytes(
			accountSigner,
			frame,
			variable,
		)
	}

	if signer.IsPublicKey() {
		pubKey := accountSigner.PublicKey()
		return app.assignableBuilder.Create().
			WithSignerPublicKey(pubKey).
			Now()
	}

	if signer.IsVoteVerify() {
		voteVerify := signer.VoteVerify()
		return app.executeVoteVerify(
			frame,
			voteVerify,
		)
	}

	if signer.IsSignatureVerify() {
		signatureVerify := signer.SignatureVerify()
		return app.executeSignatureVerify(
			accountSigner,
			frame,
			signatureVerify,
		)
	}

	if signer.IsGenerateSignerPublicKeys() {
		amount := signer.GenerateSignerPublicKeys()
		return app.executeGenerateSignerPublicKeys(
			amount,
		)
	}

	if signer.IsHashPublicKeys() {
		variable := signer.HashPublicKeys()
		return app.executeHashPublicKeys(
			frame,
			variable,
		)
	}

	if signer.IsVote() {
		vote := signer.Vote()
		return app.executeVote(
			accountSigner,
			frame,
			vote,
		)
	}

	sign := signer.Sign()
	return app.executeSign(
		accountSigner,
		frame,
		sign,
	)
}

func (app *application) executeSign(
	accountSigner account_signers.Signer,
	frame stacks.Frame,
	msgVariable string,
) (stacks.Assignable, error) {
	msg, err := frame.FetchBytes(msgVariable)
	if err != nil {
		return nil, err
	}

	sig, err := accountSigner.Sign(msg)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithSignature(sig).
		Now()
}

func (app *application) executeVote(
	accountSigner account_signers.Signer,
	frame stacks.Frame,
	signerVote layers.Vote,
) (stacks.Assignable, error) {
	ringVariable := signerVote.Ring()
	ring, err := frame.FetchSignerPublicKeys(ringVariable)
	if err != nil {
		return nil, err
	}

	msgRef := signerVote.Message()
	msg, err := frame.FetchBytes(msgRef)
	if err != nil {
		return nil, err
	}

	vote, err := accountSigner.Vote(msg, ring)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithVote(vote).
		Now()
}

func (app *application) executeGenerateSignerPublicKeys(
	amount uint,
) (stacks.Assignable, error) {
	list := []signers.PublicKey{}
	castedAmount := int(amount)
	for i := 0; i < castedAmount; i++ {
		signer := app.signerFactory.Create()
		pubKey := signer.PublicKey()
		list = append(list, pubKey)
	}

	return app.assignableBuilder.Create().
		WithSignerPublicKeys(list).
		Now()
}

func (app *application) executeHashPublicKeys(
	frame stacks.Frame,
	variable string,
) (stacks.Assignable, error) {
	pubKeys, err := frame.FetchSignerPublicKeys(variable)
	if err != nil {
		return nil, err
	}

	hashList := []hash.Hash{}
	for _, onePubKey := range pubKeys {
		bytes, err := onePubKey.Bytes()
		if err != nil {
			return nil, err
		}

		pHash, err := app.hashAdapter.FromBytes(bytes)
		if err != nil {
			return nil, err
		}

		hashList = append(hashList, *pHash)
	}

	return app.assignableBuilder.Create().
		WithHashList(hashList).
		Now()
}

func (app *application) executeVoteVerify(
	frame stacks.Frame,
	voteVerify layers.VoteVerify,
) (stacks.Assignable, error) {
	voteVariable := voteVerify.Vote()
	vote, err := frame.FetchVote(voteVariable)
	if err != nil {
		return nil, err
	}

	hashedRingVariable := voteVerify.HashedRing()
	hashedRing, err := frame.FetchHashList(hashedRingVariable)
	if err != nil {
		return nil, err
	}

	msgRef := voteVerify.Message()
	msg, err := frame.FetchBytes(msgRef)
	if err != nil {
		return nil, err
	}

	isValid, err := app.signerVoteAdapter.ToVerification(vote, msg, hashedRing)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithBool(isValid).
		Now()
}

func (app *application) executeSignatureVerify(
	accountSigner account_signers.Signer,
	frame stacks.Frame,
	signatureVerify layers.SignatureVerify,
) (stacks.Assignable, error) {
	sigVariable := signatureVerify.Signature()
	sig, err := frame.FetchSignature(sigVariable)
	if err != nil {
		return nil, err
	}

	msgRef := signatureVerify.Message()
	msg, err := frame.FetchBytes(msgRef)
	if err != nil {
		return nil, err
	}

	sigPubKey, err := sig.PublicKey(msg)
	if err != nil {
		return nil, err
	}

	isValid := accountSigner.PublicKey().Equals(sigPubKey) && sig.Verify()
	return app.assignableBuilder.Create().
		WithBool(isValid).
		Now()
}

func (app *application) executeSignerBytes(
	accountSigner account_signers.Signer,
	frame stacks.Frame,
	variable string,
) (stacks.Assignable, error) {
	assignable, err := frame.Fetch(variable)
	if err != nil {
		return nil, err
	}

	isValid := false
	builder := app.assignableBuilder.Create()
	if assignable.IsSignature() {
		data, err := assignable.Signature().Bytes()
		if err != nil {
			return nil, err
		}

		isValid = true
		builder.WithBytes(data)
	}

	if assignable.IsVote() {
		data, err := assignable.Vote().Bytes()
		if err != nil {
			return nil, err
		}

		isValid = true
		builder.WithBytes(data)
	}

	if assignable.IsSignerPublicKey() {
		data, err := assignable.SignerPublicKey().Bytes()
		if err != nil {
			return nil, err
		}

		isValid = true
		builder.WithBytes(data)
	}

	if !isValid {
		str := fmt.Sprintf("the variable (name: %s) does NOT hold a value that can be converted to bytes (signature, vote, publicKeys)", variable)
		return nil, errors.New(str)
	}

	return builder.Now()
}

func (app *application) executeEncryptor(
	accountEncryptor account_encryptors.Encryptor,
	frame stacks.Frame,
	encryptor layers.Encryptor,
) (stacks.Assignable, error) {
	if encryptor.IsDecrypt() {
		reference := encryptor.Decrypt()
		return app.executeDecrypt(
			accountEncryptor,
			frame,
			reference,
		)
	}

	if encryptor.IsEncrypt() {
		reference := encryptor.Encrypt()
		return app.executeEncrypt(
			accountEncryptor.Public(),
			frame,
			reference,
		)
	}

	pubKey := accountEncryptor.Public()
	return app.assignableBuilder.Create().
		WithEncryptorPublicKey(pubKey).
		Now()
}

func (app *application) executeDecrypt(
	encryptor account_encryptors.Encryptor,
	frame stacks.Frame,
	variable string,
) (stacks.Assignable, error) {
	cipher, err := frame.FetchBytes(variable)
	if err != nil {
		return nil, err
	}

	data, err := encryptor.Decrypt(cipher)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithBytes(data).
		Now()
}

func (app *application) executeEncrypt(
	pubKey account_encryptors.PublicKey,
	frame stacks.Frame,
	variable string,
) (stacks.Assignable, error) {
	msg, err := frame.FetchBytes(variable)
	if err != nil {
		return nil, err
	}

	cipher, err := pubKey.Encrypt(msg)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithBytes(cipher).
		Now()
}

func (app *application) executeBytes(frame stacks.Frame, bytesIns layers.Bytes) (stacks.Assignable, error) {
	if bytesIns.IsJoin() {
		variables := bytesIns.Join()
		return app.executeJoin(frame, variables)
	}

	if bytesIns.IsCompare() {
		variables := bytesIns.Join()
		return app.executeCompare(frame, variables)
	}

	variable := bytesIns.HashBytes()
	return app.executeHashBytes(frame, variable)
}

func (app *application) executeJoin(frame stacks.Frame, variables []string) (stacks.Assignable, error) {
	output := []byte{}
	for _, oneVariable := range variables {
		data, err := frame.FetchBytes(oneVariable)
		if err != nil {
			return nil, err
		}

		output = append(output, data...)
	}

	return app.assignableBuilder.Create().
		WithBytes(output).
		Now()
}

func (app *application) executeCompare(frame stacks.Frame, variables []string) (stacks.Assignable, error) {
	boolValue := true
	var lastBytes []byte
	for _, oneVariable := range variables {
		data, err := frame.FetchBytes(oneVariable)
		if err != nil {
			return nil, err
		}

		if lastBytes == nil {
			lastBytes = data
			continue
		}

		if !bytes.Equal(lastBytes, data) {
			boolValue = false
			break
		}
	}

	return app.assignableBuilder.Create().
		WithBool(boolValue).
		Now()
}

func (app *application) executeHashBytes(frame stacks.Frame, variable string) (stacks.Assignable, error) {
	data, err := frame.FetchBytes(variable)
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithHash(*pHash).
		Now()
}
