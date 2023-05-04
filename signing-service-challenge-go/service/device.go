package service

import (
	"log"

	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"

	"github.com/esnchez/coding-challenges/signing-service-challenge/crypto"
	"github.com/esnchez/coding-challenges/signing-service-challenge/persistence"
)

type SignatureService struct {
	store persistence.Repository
}

// Factory for the signature service
func NewSignatureService(repo persistence.Repository) *SignatureService {

	return &SignatureService{
		store: repo,
	}
}

// CreateSignatureDevice creates a new signature device and saves it to the repository injected
// returns the signature device created and stored
func (ss *SignatureService) CreateSignatureDevice(algorithm string, label string) (*domain.SigDevice, error) {

	signer, err := validateAndGetSigner(algorithm)
	if err != nil {
		return nil, err
	}

	dev, err := domain.NewSigDevice(signer, label)
	if err != nil {
		return nil, err
	}

	if err := ss.store.Save(dev); err != nil {
		log.Println("[ERROR] [DB] Error saving device in the repository")
		return nil, err
	}

	return dev, nil
}

// SignTransaction retrieves a device from the repository and uses it to sign some data.
// returns the final signature and the composed data that was signed
func (ss *SignatureService) SignTransaction(id uuid.UUID, data string) ([]byte, []byte, error) {

	d, err := validateDataToBeSigned(data)
	if err != nil {
		return nil, nil, err
	}

	dev, err := ss.store.Get(id)
	if err != nil {
		log.Println("[ERROR] [DB] Error retrieving device from repository")
		return nil, nil, err
	}

	signature, signedData, err := dev.Sign(d)
	if err != nil {
		return nil, nil, err
	}

	return signature, signedData, nil
}

func (ss *SignatureService) GetAllDevices() ([]*domain.SigDevice, error) {

	devList, err := ss.store.GetAll()
	if err != nil {
		log.Println("[ERROR] [DB] Error retrieving devices from repository")
		return nil, err
	}

	return devList, nil
}

// validateAndGetSigner ensures the algorithm requested by users is valid and returns a signer for creating the device
// This function will return an error if the algorithm does not match ECC or RSA. Extensible function for other implementations.
func validateAndGetSigner(algorithm string) (crypto.Signer, error) {

	switch algorithm {
	case "ECC":
		signer, err := crypto.NewECCSigner()
		if err != nil {
			return nil, err
		}
		return signer, nil
	case "RSA":
		signer, err := crypto.NewRSASigner()
		if err != nil {
			return nil, err
		}
		return signer, nil
	default:
		return nil, domain.ErrInvalidAlgorithm
	}
}

func validateDataToBeSigned(data string) ([]byte, error) {

	if data != "" {
		return []byte(data), nil
	}

	return nil, domain.ErrInvalidDataToBeSigned
}
