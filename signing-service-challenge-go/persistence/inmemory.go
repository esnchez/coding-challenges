package persistence

import (
	"sync"

	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

// TODO: in-memory persistence ...
type MemStore struct {
	mu      sync.RWMutex
	devices map[uuid.UUID]domain.SigDevice
}

func NewMemStore() *MemStore {
	return &MemStore{
		devices: make(map[uuid.UUID]domain.SigDevice),
	}
}

func (s *MemStore) Save(sd *domain.SigDevice) error {

	return nil
}

func (s *MemStore) Get(id uuid.UUID) (*domain.SigDevice, error) {

	return nil, nil
}
