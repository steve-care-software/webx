package layers

import (
	"reflect"
	"testing"
)

func TestSignatureVerify_Success(t *testing.T) {
	signature := "mySignature"
	message := "myMessage"
	signatureVerify := NewSignatureVerifyForTests(signature, message)
	retSignature := signatureVerify.Signature()
	if signature != retSignature {
		t.Errorf("the vote was expected to be '%s', '%s returned'", signature, retSignature)
		return
	}

	retMessage := signatureVerify.Message()
	if !reflect.DeepEqual(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestSignatureVerify_withoutSignature_returnsError(t *testing.T) {
	message := "myMessage"
	_, err := NewSignatureVerifyBuilder().Create().WithMessage(message).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestSignatureVerify_withoutMessage_returnsError(t *testing.T) {
	signature := "mySignature"
	_, err := NewSignatureVerifyBuilder().Create().WithSignature(signature).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
