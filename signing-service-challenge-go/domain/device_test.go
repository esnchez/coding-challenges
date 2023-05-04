package domain_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/esnchez/coding-challenges/signing-service-challenge/crypto"
	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
)

func Test_CreateNewDevice(t *testing.T) {

	eccSigner, err := crypto.NewECCSigner()
	if err != nil {
		t.Fatal(err)
	}

	label := "custom_label"

	device, err := domain.NewSigDevice(eccSigner, label)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Signature Device created with ID: %s", device.ID)
	assert.Equal(t, label, device.Label)
	assert.Equal(t, 0, device.GetCounter())
}

func Test_SignWithECCDeviceAndVerify(t *testing.T) {

	eccSigner, err := crypto.NewECCSigner()
	if err != nil {
		t.Fatal(err)
	}

	label := "custom_label"
	messageToSign := []byte("Very important message")

	device, err := domain.NewSigDevice(eccSigner, label)
	if err != nil {
		t.Fatal(err)
	}

	signature, signedData, err := device.Sign(messageToSign)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Signature created: ", string(signature))
	fmt.Println("Signed_data created: ", string(signedData))

	assert.Equal(t, 1, device.GetCounter())

	is, err := eccSigner.Verify(signedData, signature)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, is)
}

func Test_Sign2TimesWithECCDevice(t *testing.T) {

	eccSigner, err := crypto.NewECCSigner()
	if err != nil {
		t.Fatal(err)
	}

	label := "custom_label"
	messageToSign := []byte("Very important message")

	device, err := domain.NewSigDevice(eccSigner, label)
	if err != nil {
		t.Fatal(err)
	}

	signature, sigData, err := device.Sign(messageToSign)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Signature created: ", string(signature))
	fmt.Println("Signed_data created: ", string(sigData))

	assert.Equal(t, 1, device.GetCounter())

	signature2, sigData2, err := device.Sign(messageToSign)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Signature2 created: ", string(signature2))
	fmt.Println("Signed data2 created: ", string(sigData2))

	assert.Equal(t, 2, device.GetCounter())
}


