package persistence

import (
	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Save(*domain.SigDevice) error
	Get(uuid.UUID) (*domain.SigDevice, error)
}
