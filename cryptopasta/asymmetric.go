package cryptopasta

import (
	"crypto/rsa"
	"crypto"
	"crypto/rand"
	"store/util"
)

// GenerateKeys generate private and public keys
func GenerateKeys(size int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	return privateKey, &privateKey.PublicKey, err
}

// EncryptOAEP encrypts the given message with RSA-OAEP.
func EncryptOAEP(publicKey *rsa.PublicKey, msg interface{}) (out []byte, err error) {
	msgBytes, err := util.StructToBytes(msg)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(crypto.SHA256.New(), rand.Reader, publicKey, msgBytes, nil)
}

// DecryptOAEP decrypts ciphertext using RSA-OAEP.
func DecryptOAEP(privateKey *rsa.PrivateKey, cipherText []byte, msg interface{}) (err error) {
	msgBytes, err := rsa.DecryptOAEP(crypto.SHA256.New(), rand.Reader, privateKey, cipherText, nil)
	if err != nil {
		return err
	}
	return util.BytesToStruct(msgBytes, msg)
}
