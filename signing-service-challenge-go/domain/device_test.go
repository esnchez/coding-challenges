package domain_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/esnchez/coding-challenges/signing-service-challenge/crypto"
	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
)

func Test_CreateNewDevice(t *testing.T) {

	type testCase struct {
		test        string
		label       string
		signer      crypto.Signer
		expectedErr error
	}

	eccSigner := &crypto.ECCSigner{}
	rsaSigner := &crypto.RSASigner{}

	testCases := []testCase{
		{
			test:        "Creating a new signature device with ECC signer injected and label",
			label:       "custom",
			signer:      eccSigner,
			expectedErr: nil,
		},
		{
			test:        "Creating a new signature device with RSA signer injected, no label",
			signer:      rsaSigner,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			sigDev, err := domain.NewSigDevice(tc.signer, tc.label)
			if err != nil {
				t.Fatal(err)
			}

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			assert.Equal(t, tc.label, sigDev.Label)
			assert.Equal(t, 0, sigDev.GetCounter())
		})
	}
}

func Test_SignWithDeviceAndVerify(t *testing.T) {

	type testCase struct {
		test        string
		label       string
		signer      crypto.Signer
		message     []byte
		expectedErr error
	}

	eccSigner, err := crypto.NewECCSigner()
	if err != nil {
		t.Fatal(err)
	}
	rsaSigner, err := crypto.NewRSASigner()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []testCase{
		{
			test:        "Sign message with a signature device with ECC signer",
			label:       "custom",
			signer:      eccSigner,
			message:     []byte("Very important message"),
			expectedErr: nil,
		},
		{
			test:        "Sign message with a signature device with RSA signer",
			signer:      rsaSigner,
			message:     []byte("Very important message"),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			sigDev, err := domain.NewSigDevice(tc.signer, tc.label)
			if err != nil {
				t.Fatal(err)
			}

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			assert.Equal(t, 0, sigDev.GetCounter())

			signature, signedData, err := sigDev.Sign(tc.message)
			if err != nil {
				t.Fatal(err)
			}

			is, err := tc.signer.Verify(signedData, signature)
			if err != nil {
				t.Fatal(err)
			}

			//checking counter get to 1 and verifying the signature
			assert.Equal(t, 1, sigDev.GetCounter())
			assert.Equal(t, true, is)
		})
	}

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

	_, _, err = device.Sign(messageToSign)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, device.GetCounter())

	_, _, err = device.Sign(messageToSign)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, device.GetCounter())
}
