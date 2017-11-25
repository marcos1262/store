package main

import (
	"net/rpc"
	"net"
	"store/store"
	"store/server/RPC"
	"log"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":54321")
	store.CheckMortalErr(err)

	listener, err := net.ListenTCP("tcp", addr)
	store.CheckMortalErr(err)

	log.Println("Registering RPC CRUDs")
	product := new(RPC.RPC_product)
	rpc.Register(product)

	user := new(RPC.RPC_user)
	rpc.Register(user)

	log.Println("Listening to clients on port 54321...")
	rpc.Accept(listener)
}
