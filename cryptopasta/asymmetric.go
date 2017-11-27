package cryptopasta

import (
	"crypto/rsa"
	"crypto"
	"crypto/rand"
)

// EncryptOAEP encrypts the given message with RSA-OAEP.
func EncryptOAEP(pub *rsa.PublicKey, msg []byte) (out []byte, err error) {
	return rsa.EncryptOAEP(crypto.SHA256.New(), rand.Reader, pub, msg, nil)
}

// DecryptOAEP decrypts ciphertext using RSA-OAEP.
func DecryptOAEP(priv *rsa.PrivateKey, ciphertext []byte) (msg []byte, err error) {
	return rsa.DecryptOAEP(crypto.SHA256.New(), rand.Reader, priv, ciphertext, nil)
}
