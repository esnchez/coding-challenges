package crypto

import (
	"errors"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
	Verify(dataToVerify []byte, sig []byte) (bool, error)
}

var ErrSigningFailed = errors.New("Error ocurred during signing process")
var ErrSignerCreationFailed = errors.New("Error ocurred during signing process")
var ErrVerifyDecodeFailed = errors.New("Error ocurred during decoding while verifying process")
var ErrVerifyRSASignatureFailed = errors.New("Verifying RSA signature process failed")

//TODO: implement RSA and ECDSA signing ...
//UPDATABLE: If another algorithm is added to the system, build the logic here ...
