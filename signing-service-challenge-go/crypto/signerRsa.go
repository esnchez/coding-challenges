package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// RSA
type RSASigner struct {
	PubKey  *rsa.PublicKey
	PrivKey *rsa.PrivateKey
}

func NewRSASigner() (*RSASigner, error) {

	rsaGen := new(RSAGenerator)
	rsaKeyPair, err := rsaGen.Generate()
	if err != nil {
		return nil, err
	}

	return &RSASigner{
		PubKey:  rsaKeyPair.Public,
		PrivKey: rsaKeyPair.Private,
	}, nil
}

func (rs *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {

	msgHash := sha256.New()
	_, err := msgHash.Write(dataToBeSigned)
	if err != nil {
		return nil, err
	}
	msgHashSum := msgHash.Sum(nil)
	fmt.Println("msgHashSum", msgHashSum)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	sig, err := rsa.SignPSS(rand.Reader, rs.PrivKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		return nil, ErrSigningFailed
	}

	signature := []byte(base64.StdEncoding.EncodeToString(sig))
	fmt.Println("[DEBUG] Printing signature after base64 encoding:", string(signature))
	return signature, nil
}

func (rs *RSASigner) Verify(dataToVerify []byte, sig []byte) (bool, error) {

	//hashing data
	msgHash := sha256.New()
	_, err := msgHash.Write(dataToVerify)
	if err != nil {
		return false, err
	}
	msgHashSum := msgHash.Sum(nil)
	fmt.Println("msgHashSum", msgHashSum)

	//Base64 Decode signature step
	signature, err := base64.StdEncoding.DecodeString(string(sig))
	if err != nil {
		return false, ErrVerifyDecodeFailed
	}

	if err := rsa.VerifyPSS(rs.PubKey, crypto.SHA256, msgHashSum, signature, nil); err != nil {
		return false, ErrVerifyRSASignatureFailed
	}

	return true, nil
}
