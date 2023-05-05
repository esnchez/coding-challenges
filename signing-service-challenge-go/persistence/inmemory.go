package persistence

import (
	"fmt"
	"sync"

	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

// TODO: in-memory persistence ...
//MODIFY mutex or rwmutex?

type MemStore struct {
	mu      sync.Mutex
	devices map[uuid.UUID]*domain.SigDevice
}

func NewMemStore() *MemStore {
	return &MemStore{
		devices: make(map[uuid.UUID]*domain.SigDevice),
	}
}

func (s *MemStore) Save(sd *domain.SigDevice) error {
	if _, ok := s.devices[sd.ID]; ok {
		return ErrSaveFailure
	}

	s.mu.Lock()
	s.devices[sd.ID] = sd
	s.mu.Unlock()
	fmt.Println("Saved!", sd)
	return nil
}

func (s *MemStore) Get(id uuid.UUID) (*domain.SigDevice, error) {
	if device, ok := s.devices[id]; ok {
		fmt.Println("[DEBUD]: recovered!")
		return device, nil
	}

	return &domain.SigDevice{}, ErrNotFound
}

func (s *MemStore) GetAll() ([]*domain.SigDevice, error) {

	list := []*domain.SigDevice{}
	for _, v := range s.devices {
		
		fmt.Printf("Printing what is stored: %s\n", v.ID)
		list = append(list, v)
		fmt.Println("Printing the list", list)
	}

	return list, nil
}
