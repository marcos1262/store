package main

import (
	"net"
	"bufio"
	"store/util"
	"store/cryptopasta"
	"crypto/rsa"
	"store/model"
	"encoding/hex"
	"crypto/sha256"
	"log"
	"os"
)

func authenticate(conn net.Conn) {
	in := bufio.NewReader(conn)
	out := bufio.NewWriter(conn)

	// Send public key
	publicKeyBytes, err := util.StructToBytes(publicKey) // 1866 bytes
	util.CheckMortalErr(err)

	err = util.SendData(out, publicKeyBytes)
	util.CheckMortalErr(err)

	// Receive encrypted server's public key
	encrypted, err := util.ReceiveData(in, 768)
	util.CheckMortalErr(err)

	serverPublicKey := &rsa.PublicKey{}
	err = cryptopasta.DecryptOAEP(privateKey, encrypted, serverPublicKey)
	util.CheckMortalErr(err)

	// Send auth info encrypted
	pass := sha256.Sum256([]byte("123456"))
	user := model.User{
		Login: "admin",
		Pass:  hex.EncodeToString(pass[:]),
	}

	encrypted, err = cryptopasta.EncryptOAEP(serverPublicKey, user) // 256 bytes
	util.CheckMortalErr(err)

	err = util.SendData(out, encrypted)
	util.CheckMortalErr(err)

	// Receive encrypted session key or error (if not authenticate)
	encrypted, err = util.ReceiveData(in, 768)
	if err != nil {
		log.Fatal(util.ERR_NOT_AUTH)
		os.Exit(1)
	}

	err = cryptopasta.DecryptOAEP(privateKey, encrypted, sessionKey)
	util.CheckMortalErr(err)
}
