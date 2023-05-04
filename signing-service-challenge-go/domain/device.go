package domain

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sync"

	"github.com/esnchez/coding-challenges/signing-service-challenge/crypto"

	"github.com/google/uuid"
)

// TODO: signature device domain model ...

var ErrInvalidAlgorithm = errors.New("Signature device cannot be created with the specified data")
var ErrInvalidDataToBeSigned = errors.New("Data to be signed is required")
var ErrSignOperation = errors.New("Signature device failed during signing operation")

type SigDevice struct {
	ID            uuid.UUID
	signer        crypto.Signer
	lastSignature []byte
	Label         string
	mu            sync.Mutex
	counter       int
}

// Factory
func NewSigDevice(signer crypto.Signer, label string) (*SigDevice, error) {

	return &SigDevice{
		ID:     uuid.New(),
		signer: signer,
		Label:  label,
	}, nil
}

// Sign uses the signer embedded in the device for signing some data and updates its own state
func (s *SigDevice) Sign(data []byte) ([]byte, []byte, error) {

	toSignData := s.generateDataToSign(data)

	sig, err := s.signer.Sign(toSignData)
	if err != nil {
		return nil, nil, ErrSignOperation
	}

	// s.counter++
	// s.lastSignature = sig
	go s.updateDeviceState(sig)

	return sig, toSignData, nil
}

func (s *SigDevice) updateDeviceState(sig []byte) {
	s.mu.Lock()
	s.counter++
	s.lastSignature = sig
	s.mu.Unlock()
}

// generateDataToSign is a private function that applies some business rules to our data previously to being signed
func (s *SigDevice) generateDataToSign(data []byte) []byte {

	if s.counter == 0 {
		encID := base64.StdEncoding.EncodeToString([]byte(s.ID.String()))
		fmt.Println("[DEBUG]: Printing lastSignature (when counter equals 0), which is base64 encoded signature device ID:", encID)

		return []byte(fmt.Sprintf("%d_%s_%s", s.counter, data, encID))
	}

	return []byte(fmt.Sprintf("%d_%s_%s", s.counter, data, s.lastSignature))
}

func (s *SigDevice) setCounter(newCounter int) {
	s.counter = newCounter
}

func (s *SigDevice) GetCounter() int {
	return s.counter
}
