package domain

import "github.com/google/uuid"

// TODO: signature device domain model ...

type SigDevice struct {
	id         uuid.UUID
	// pubKey     
	// privKey    
	label      string
	sigCounter int
}

func NewSigDevice(label string) *SigDevice {
	return &SigDevice{
		id:    uuid.New(),
		label: label,
	}
}

func (s *SigDevice) SetCounter(newCounter int) {
	s.sigCounter = newCounter
}

func (s *SigDevice) GetID() uuid.UUID {
	return s.id
}

func (s *SigDevice) GetLabel() string {
	return s.label
}
