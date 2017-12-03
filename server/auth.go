package main

import (
	"net"
	"store/util"
	"store/model"
	"bufio"
	"store/cryptopasta"
	"crypto/rsa"
	"store/server/DAO"
	"crypto/sha256"
	"time"
	"log"
)

// Exchanges cryptographic messages with client to verify his authenticity
func authenticateClient(conn net.Conn) {
	client := &model.Client{
		Conn:      conn,
		In:        bufio.NewReader(conn),
		Out:       bufio.NewWriter(conn),
		PublicKey: &rsa.PublicKey{},
	}

	// Receive client's public key
	publicKeyBytes, err := util.ReceiveDataLine(client.In)
	util.CheckErr(err)

	err = util.BytesToStruct(publicKeyBytes, client.PublicKey)
	util.CheckErr(err)

	// Send encrypted server's public key
	encrypted, err := cryptopasta.EncryptOAEP(client.PublicKey, publicKey) // 768 bytes
	util.CheckErr(err)

	err = util.SendData(client.Out, encrypted)
	util.CheckErr(err)

	// Receive encrypted client's auth info
	encrypted, err = util.ReceiveData(client.In, 256)
	util.CheckErr(err)

	user := &model.User{}
	err = cryptopasta.DecryptOAEP(privateKey, encrypted, user)
	util.CheckErr(err)

	// Authenticate user
	_, total, err := DAO.ReadUser(user, 0, 1)
	util.CheckErr(err)
	if total < 1 {
		client.Conn.Close()
		return
	}

	// Send encrypted session key
	userBytes, err := util.StructToBytes(user)
	util.CheckErr(err)

	sessionKey := sha256.Sum256(append(userBytes, []byte(time.Now().String())...))
	client.SessionKey = &sessionKey

	encrypted, err = cryptopasta.EncryptOAEP(client.PublicKey, sessionKey) // 768 bytes
	util.CheckErr(err)

	err = util.SendData(client.Out, encrypted)
	util.CheckErr(err)

	// Save client
	saveClient(string(publicKeyBytes), client)

	go expiresSession(string(publicKeyBytes), client)
}

// Closes client's connection after some time
func expiresSession(key string, client *model.Client) {
	time.Sleep(60 * time.Second)

	log.Println("Expiring session of client", key)
	delClient(key)
	client.Conn.Close()
}
