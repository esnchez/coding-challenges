package persistence

import (
	"errors"

	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("Signature device could not be found in the repository")
var ErrSaveFailure = errors.New("Signature device failed to be added, it already exists")

//interface that describes the contract that all persistence implementations have to comply with
type Repository interface {
	Save(*domain.SigDevice) error
	Get(uuid.UUID) (*domain.SigDevice, error)
	GetAll() ([]*domain.SigDevice, error)
}
