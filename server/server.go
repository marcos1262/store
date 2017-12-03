package main

import (
	"net/rpc"
	"net"
	"store/util"
	"store/server/RPC"
	"log"
	"store/server/DAO"
	"crypto/rsa"
	"store/cryptopasta"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":54321")
	util.CheckMortalErr(err)

	listener, err := net.ListenTCP("tcp", addr)
	util.CheckMortalErr(err)

	log.Println("Initializing Database")
	DAO.InitializeDB()
	defer DAO.CloseDB()

	log.Println("Generating keys for cryptography")
	privateKey, publicKey, err = cryptopasta.GenerateKeys(2048)
	util.CheckMortalErr(err)

	log.Println("Registering RPC CRUDs")
	rpc.Register(new(RPC.RPC_product))
	rpc.Register(new(RPC.RPC_user))

	log.Println("Managing clients...")
	go manageClients()

	log.Println("Listening to clients on port 54321...")
	for {
		conn, err := listener.Accept()
		if util.CheckErr(err) {
			continue
		}
		go func() {
			authenticateClient(conn)
			rpc.ServeConn(conn)
		}()
	}
}
