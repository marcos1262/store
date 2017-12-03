package cryptopasta

import (
	"crypto/rsa"
	"store/util"
)

type EncryptedMsg struct {
	PublicKey []byte
	Msg       []byte
}

func NewEncryptedMsg(
	publicKey *rsa.PublicKey,
	otherPublicKey *rsa.PublicKey,
	sessionKey *[32]byte,
	msg interface{}) (*EncryptedMsg, error) {

	publicKeyBytes, err := util.StructToBytes(publicKey)
	if err != nil {
		return nil, err
	}
	publicKeyEncrypted, err := EncryptOAEP(otherPublicKey, publicKeyBytes)
	if err != nil {
		return nil, err
	}

	dataBytes, err := util.StructToBytes(msg)
	if err != nil {
		return nil, err
	}
	msgEncrypted, err := EncryptAES(dataBytes, sessionKey)
	if err != nil {
		return nil, err
	}

	return &EncryptedMsg{
		PublicKey: publicKeyEncrypted,
		Msg:       msgEncrypted,
	}, nil
}

//func DecryptMsg(privateKey *rsa.PrivateKey, sessionKey ){
//	decrypted, err := cryptopasta.DecryptAES(encrypted, SessionKey)
//	println(string(decrypted))
//}