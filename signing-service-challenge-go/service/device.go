package service

import "github.com/esnchez/coding-challenges/signing-service-challenge/persistence"

type SignatureService struct {
	store persistence.Repository
}

func NewSignatureService(repo persistence.Repository) *SignatureService {
	return &SignatureService{
		store: repo,
	}
}