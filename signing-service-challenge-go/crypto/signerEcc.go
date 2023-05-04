package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

// ECDSA
type ECCSigner struct {
	PubKey  *ecdsa.PublicKey
	PrivKey *ecdsa.PrivateKey
}

func NewECCSigner() (*ECCSigner, error) {

	eccGen := new(ECCGenerator)
	eccKeyPair, err := eccGen.Generate()
	if err != nil {
		return nil, ErrSignerCreationFailed
	}

	return &ECCSigner{
		PubKey:  eccKeyPair.Public,
		PrivKey: eccKeyPair.Private,
	}, nil
}

func (ecc *ECCSigner) Sign(dataToBeSigned []byte) ([]byte, error) {

	//Add a hashing operation to the data before signing?
	// hash := sha256.Sum256(dataToBeSigned)
	// fmt.Println("[DEBUG] printing dataToBeSigned AFTER hashing: ", hash)
	msgHash := sha256.New()
	_, err := msgHash.Write(dataToBeSigned)
	if err != nil {
		return nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, ecc.PrivKey, msgHashSum)
	if err != nil {
		return nil, ErrSigningFailed
	}

	sig := append(r.Bytes(), s.Bytes()...)

	//Base64 encode step before returning signature
	signature := []byte(base64.StdEncoding.EncodeToString(sig))
	fmt.Println("[DEBUG] Printing signature after base64 encoding:", string(signature))

	return signature, nil
}

func (ecc *ECCSigner) Verify(dataToVerify []byte, sig []byte) (bool, error) {

	msgHash := sha256.New()
	_, err := msgHash.Write(dataToVerify)
	if err != nil {
		return false, err
	}
	msgHashSum := msgHash.Sum(nil)

	//Base64 decode step
	signature, err := base64.StdEncoding.DecodeString(string(sig))
	if err != nil {
		return false, ErrVerifyDecodeFailed
	}

	r := big.Int{}
	s := big.Int{}
	len := len(signature)

	r.SetBytes(signature[:(len / 2)])
	s.SetBytes(signature[(len / 2):])

	return ecdsa.Verify(ecc.PubKey, msgHashSum, &r, &s), nil
}
