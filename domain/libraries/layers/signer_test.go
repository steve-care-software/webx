package layers

import (
	"reflect"
	"testing"
)

func TestSigner_withSign_Success(t *testing.T) {
	sign := "mySign"
	signer := NewSignerWithSignForTests(sign)

	if !signer.IsSign() {
		t.Errorf("the signer was expected to contain a sign")
		return
	}

	if signer.IsVote() {
		t.Errorf("the signer was expected to NOT contain a vote")
		return
	}

	if signer.IsGenerateSignerPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a generateSignerPublicKeys")
		return
	}

	if signer.IsHashPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a hashPublicKeys")
		return
	}

	if signer.IsVoteVerify() {
		t.Errorf("the signer was expected to NOT contain a voteVerify")
		return
	}

	if signer.IsSignatureVerify() {
		t.Errorf("the signer was expected to NOT contain a signatureVerify")
		return
	}

	if signer.IsBytes() {
		t.Errorf("the signer was expected to NOT contain a bytes")
		return
	}

	if signer.IsPublicKey() {
		t.Errorf("the signer was expected to NOT contain a publicKey")
		return
	}

	retSign := signer.Sign()
	if !reflect.DeepEqual(sign, retSign) {
		t.Errorf("the returned sign is invalid")
		return
	}
}

func TestSigner_withGenerateSignerPublicKey_Success(t *testing.T) {
	amount := uint(32)
	signer := NewSignerWithGenerateSignerPublicKeyForTests(amount)

	if signer.IsSign() {
		t.Errorf("the signer was expected to NOT contain a sign")
		return
	}

	if signer.IsVote() {
		t.Errorf("the signer was expected to NOT contain a vote")
		return
	}

	if !signer.IsGenerateSignerPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a generateSignerPublicKeys")
		return
	}

	if signer.IsHashPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a hashPublicKeys")
		return
	}

	if signer.IsVoteVerify() {
		t.Errorf("the signer was expected to NOT contain a voteVerify")
		return
	}

	if signer.IsSignatureVerify() {
		t.Errorf("the signer was expected to NOT contain a signatureVerify")
		return
	}

	if signer.IsBytes() {
		t.Errorf("the signer was expected to NOT contain a bytes")
		return
	}

	if signer.IsPublicKey() {
		t.Errorf("the signer was expected to NOT contain a publicKey")
		return
	}

	retAmount := signer.GenerateSignerPublicKeys()
	if retAmount != amount {
		t.Errorf("the returned generateSignerPublicKeys is invalid, %d expected, %d returned", amount, retAmount)
		return
	}
}

func TestSigner_withHashPublicKeys_Success(t *testing.T) {
	hashPublicKeys := "myKeys"
	signer := NewSignerWithHashPublicKeysForTests(hashPublicKeys)

	if signer.IsSign() {
		t.Errorf("the signer was expected to NOT contain a sign")
		return
	}

	if signer.IsVote() {
		t.Errorf("the signer was expected to NOT contain a vote")
		return
	}

	if signer.IsGenerateSignerPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a generateSignerPublicKeys")
		return
	}

	if !signer.IsHashPublicKeys() {
		t.Errorf("the signer was expected to contain a hashPublicKeys")
		return
	}

	if signer.IsVoteVerify() {
		t.Errorf("the signer was expected to NOT contain a voteVerify")
		return
	}

	if signer.IsSignatureVerify() {
		t.Errorf("the signer was expected to NOT contain a signatureVerify")
		return
	}

	if signer.IsBytes() {
		t.Errorf("the signer was expected to NOT contain a bytes")
		return
	}

	if signer.IsPublicKey() {
		t.Errorf("the signer was expected to NOT contain a publicKey")
		return
	}

	retKeys := signer.HashPublicKeys()
	if hashPublicKeys != retKeys {
		t.Errorf("the returned hashPublicKeys is invalid, '%s' expected, '%s' returned", hashPublicKeys, retKeys)
		return
	}
}

func TestSigner_withVoteVerifty_Success(t *testing.T) {
	vote := "myVote"
	message := "myMessage"
	hashedRing := "myHashedRingVariable"
	voteVerify := NewVoteVerifyForTests(vote, message, hashedRing)
	signer := NewSignerWithVoteVerifyForTests(voteVerify)

	if signer.IsSign() {
		t.Errorf("the signer was expected to NOT contain a sign")
		return
	}

	if signer.IsVote() {
		t.Errorf("the signer was expected to NOT contain a vote")
		return
	}

	if signer.IsGenerateSignerPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a generateSignerPublicKeys")
		return
	}

	if signer.IsHashPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a hashPublicKeys")
		return
	}

	if !signer.IsVoteVerify() {
		t.Errorf("the signer was expected to contain a voteVerify")
		return
	}

	if signer.IsSignatureVerify() {
		t.Errorf("the signer was expected to NOT contain a signatureVerify")
		return
	}

	if signer.IsBytes() {
		t.Errorf("the signer was expected to NOT contain a bytes")
		return
	}

	if signer.IsPublicKey() {
		t.Errorf("the signer was expected to NOT contain a publicKey")
		return
	}

	retVoteVerify := signer.VoteVerify()
	if !reflect.DeepEqual(voteVerify, retVoteVerify) {
		t.Errorf("the returned voteVerify is invalid")
		return
	}
}

func TestSigner_withSignatureVerifty_Success(t *testing.T) {
	signature := "mySignature"
	message := "myMessage"
	signatureVerify := NewSignatureVerifyForTests(signature, message)
	signer := NewSignerWithSignatureVerifyForTests(signatureVerify)

	if signer.IsSign() {
		t.Errorf("the signer was expected to NOT contain a sign")
		return
	}

	if signer.IsVote() {
		t.Errorf("the signer was expected to NOT contain a vote")
		return
	}

	if signer.IsGenerateSignerPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a generateSignerPublicKeys")
		return
	}

	if signer.IsHashPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a hashPublicKeys")
		return
	}

	if signer.IsVoteVerify() {
		t.Errorf("the signer was expected to NOT contain a voteVerify")
		return
	}

	if !signer.IsSignatureVerify() {
		t.Errorf("the signer was expected to contain a signatureVerify")
		return
	}

	if signer.IsBytes() {
		t.Errorf("the signer was expected to NOT contain a bytes")
		return
	}

	if signer.IsPublicKey() {
		t.Errorf("the signer was expected to NOT contain a publicKey")
		return
	}

	retSignatureVerify := signer.SignatureVerify()
	if !reflect.DeepEqual(signatureVerify, retSignatureVerify) {
		t.Errorf("the returned signatureVerify is invalid")
		return
	}
}

func TestSigner_withBytes_Success(t *testing.T) {
	str := "thisIsBytes"
	signer := NewSignerWithBytesForTests(str)

	if signer.IsSign() {
		t.Errorf("the signer was expected to NOT contain a sign")
		return
	}

	if signer.IsVote() {
		t.Errorf("the signer was expected to NOT contain a vote")
		return
	}

	if signer.IsGenerateSignerPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a generateSignerPublicKeys")
		return
	}

	if signer.IsHashPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a hashPublicKeys")
		return
	}

	if signer.IsVoteVerify() {
		t.Errorf("the signer was expected to NOT contain a voteVerify")
		return
	}

	if signer.IsSignatureVerify() {
		t.Errorf("the signer was expected to NOT contain a signatureVerify")
		return
	}

	if !signer.IsBytes() {
		t.Errorf("the signer was expected to contain a bytes")
		return
	}

	if signer.IsPublicKey() {
		t.Errorf("the signer was expected to NOT contain a publicKey")
		return
	}

	retBytes := signer.Bytes()
	if str != retBytes {
		t.Errorf("the bytes variable was expected to be '%s', '%s' returned", str, retBytes)
		return
	}
}

func TestSigner_isPublicKey_Success(t *testing.T) {
	signer := NewSignerWithPublicKeyForTests()

	if signer.IsSign() {
		t.Errorf("the signer was expected to NOT contain a sign")
		return
	}

	if signer.IsVote() {
		t.Errorf("the signer was expected to NOT contain a vote")
		return
	}

	if signer.IsGenerateSignerPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a generateSignerPublicKeys")
		return
	}

	if signer.IsHashPublicKeys() {
		t.Errorf("the signer was expected to NOT contain a hashPublicKeys")
		return
	}

	if signer.IsVoteVerify() {
		t.Errorf("the signer was expected to NOT contain a voteVerify")
		return
	}

	if signer.IsSignatureVerify() {
		t.Errorf("the signer was expected to NOT contain a signatureVerify")
		return
	}

	if signer.IsBytes() {
		t.Errorf("the signer was expected to NOT contain a bytes")
		return
	}

	if !signer.IsPublicKey() {
		t.Errorf("the signer was expected to contain a publicKey")
		return
	}
}

func TestSigner_withoutParam_returnsError(t *testing.T) {
	_, err := NewSignerBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
