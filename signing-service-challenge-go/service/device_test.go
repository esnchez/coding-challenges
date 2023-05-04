package service_test

import (
	"fmt"
	"testing"

	"github.com/esnchez/coding-challenges/signing-service-challenge/persistence"
	"github.com/esnchez/coding-challenges/signing-service-challenge/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CreateSignatureDevice(t *testing.T) {
	fmt.Println("init")

	mem := persistence.NewMemStore()

	svc := service.NewSignatureService(mem)

	alg := "ECC"
	label := "custom_label"

	dev, err := svc.CreateSignatureDevice(alg, label)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Signature Device created with ID: %s\n", dev.ID)
	assert.Equal(t, label, dev.Label)
	assert.Equal(t, 0, dev.GetCounter())
	fmt.Println("end")

	rec, err := mem.Get(dev.ID)
	assert.Equal(t, rec.ID, dev.ID)
	assert.Equal(t, rec.Label, dev.Label)
	assert.Equal(t, rec.GetCounter(), dev.GetCounter())
}

func Test_SignTransactionWithInvalidOrNoDevice(t *testing.T) {
	fmt.Println("init")

	mem := persistence.NewMemStore()

	svc := service.NewSignatureService(mem)

	msg := "Testing message transaction"

	_, _ , err := svc.SignTransaction(uuid.UUID{}, msg)
	if err != nil {
		assert.Equal(t, persistence.ErrNotFound, err)
	}
}

func Test_SignTransactionWithDevice(t *testing.T) {
	fmt.Println("init")

	mem := persistence.NewMemStore()
	
	svc := service.NewSignatureService(mem)

	alg := "ECC"
	label := "custom_label"

	dev, err := svc.CreateSignatureDevice(alg, label)
	if err != nil {
		t.Fatal(err)
	}

	msg := "Testing message transaction"


	_ ,_, err = svc.SignTransaction(dev.ID, msg)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1 , dev.GetCounter())

	_ ,_, err = svc.SignTransaction(dev.ID, msg)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2 , dev.GetCounter())
}

