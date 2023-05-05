package crypto_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/esnchez/coding-challenges/signing-service-challenge/crypto"
)

func Test_SignAndVerifyWithRSASigner(t *testing.T) {

	rsaSigner, err := crypto.NewRSASigner()
	if err != nil {
		t.Fatal(err)
	}

	rsaSigner2, err := crypto.NewRSASigner()
	if err != nil {
		t.Fatal(err)
	}

	messageToSign := []byte("I love pasta")

	signature, err := rsaSigner.Sign(messageToSign)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Signature created: ", string(signature))

	is, err := rsaSigner.Verify(messageToSign, signature)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, is)

	is, err = rsaSigner2.Verify(messageToSign, signature)
	if err != nil {
		assert.Equal(t, false, is)
	}

}

func Test_SignAndVerifyWithECCSigner(t *testing.T) {

	eccSigner, err := crypto.NewECCSigner()
	if err != nil {
		t.Fatal(err)
	}

	eccSigner2, err := crypto.NewECCSigner()
	if err != nil {
		t.Fatal(err)
	}

	messageToSign := []byte("I love pasta")

	signature, err := eccSigner.Sign(messageToSign)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Signature created: ", string(signature))

	is, err := eccSigner.Verify(messageToSign, signature)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, is)

	is, err = eccSigner2.Verify(messageToSign, signature)
	if err != nil {
		assert.Equal(t, false, is)

	}
}
