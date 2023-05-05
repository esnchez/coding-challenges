package service_test

import (
	"errors"
	"testing"

	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/esnchez/coding-challenges/signing-service-challenge/persistence"
	"github.com/esnchez/coding-challenges/signing-service-challenge/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CreateSignatureServiceWithDifferentAlgorithms(t *testing.T) {

	type testCase struct {
		test        string
		algorithm   string
		label       string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Create device from service with ECC algorithm and label",
			label:       "custom",
			algorithm:   "ECC",
			expectedErr: nil,
		},
		{
			test:        "Create device from service with RSA algorithm, no label",
			algorithm:   "RSA",
			expectedErr: nil,
		},
		{
			test:        "Create device from service with random algorithm",
			algorithm:   "Random",
			expectedErr: domain.ErrInvalidAlgorithm,
		},
		{
			test:        "Create device from service with no algorithm",
			expectedErr: domain.ErrInvalidAlgorithm,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {

			mem := persistence.NewMemStore()
			svc := service.NewSignatureService(mem)

			sigDev, err := svc.CreateSignatureDevice(tc.algorithm, tc.label)
			if err != nil {
				// t.Fatal(err)

				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("expected error %v, got %v", tc.expectedErr, err)
				}
				assert.Equal(t, err, tc.expectedErr)
				return
			}

			assert.Equal(t, tc.label, sigDev.Label)
			assert.Equal(t, 0, sigDev.GetCounter())

			//recovering from storage and comparing
			rec, err := mem.Get(sigDev.ID)
			assert.Equal(t, rec.ID, sigDev.ID)
			assert.Equal(t, rec.Label, sigDev.Label)
			assert.Equal(t, rec.GetCounter(), sigDev.GetCounter())

		})
	}
}

func Test_SignTransactionWithNoDevice(t *testing.T) {

	mem := persistence.NewMemStore()
	svc := service.NewSignatureService(mem)

	msg := "Testing message transaction"

	_, _, err := svc.SignTransaction(uuid.UUID{}, msg)
	if err != nil {
		assert.Equal(t, persistence.ErrNotFound, err)
	}
}

func Test_SignTransactionWithInvalidId(t *testing.T) {

	mem := persistence.NewMemStore()
	svc := service.NewSignatureService(mem)

	_, err := svc.CreateSignatureDevice("ECC", "custom")
	if err != nil {
		t.Fatal(err)
	}

	msg := "Testing message transaction"

	_, _, err = svc.SignTransaction(uuid.UUID{}, msg)
	if err != nil {
		assert.Equal(t, persistence.ErrNotFound, err)
	}
}

func Test_SignTransactionFromServiceWithDevice(t *testing.T) {
	type testCase struct {
		test        string
		algorithm   string
		label       string
		message     string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Sign from service, recovering a device created with ECC algorithm and label",
			label:       "custom",
			algorithm:   "ECC",
			message:     "Very important message",
			expectedErr: nil,
		},
		{
			test:        "Sign from service, recovering a device created with RSA algorithm, no label",
			algorithm:   "RSA",
			message:     "Very important message",
			expectedErr: nil,
		},
		{
			test:        "Sign from service, recovering a device created with ECC algorithm but No data",
			algorithm:   "ECC",
			expectedErr: domain.ErrInvalidDataToBeSigned,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {

			mem := persistence.NewMemStore()
			svc := service.NewSignatureService(mem)

			sigDev, err := svc.CreateSignatureDevice(tc.algorithm, tc.label)
			if err != nil {
				// t.Fatal(err)

				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("expected error %v, got %v", tc.expectedErr, err)
				}
				assert.Equal(t, err, tc.expectedErr)
				return
			}

			assert.Equal(t, tc.label, sigDev.Label)
			assert.Equal(t, 0, sigDev.GetCounter())

			_, _, err = svc.SignTransaction(sigDev.ID, tc.message)
			if err != nil {
				//t.Fatal(err)
				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("expected error %v, got %v", tc.expectedErr, err)
				}
				assert.Equal(t, err, tc.expectedErr)
				return
			}

			assert.Equal(t, 1, sigDev.GetCounter())

		})
	}
}
