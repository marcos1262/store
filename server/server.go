package main

import (
	"net/rpc"
	"net"
	"store/util"
	"store/server/RPC"
	"log"
	"store/server/DAO"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":54321")
	util.CheckMortalErr(err)

	listener, err := net.ListenTCP("tcp", addr)
	util.CheckMortalErr(err)

	log.Println("Initializing Database")
	DAO.InitializeDB()
	defer DAO.CloseDB()

	log.Println("Registering RPC CRUDs")
	rpc.Register(new(RPC.RPC_auth))
	rpc.Register(new(RPC.RPC_product))
	rpc.Register(new(RPC.RPC_user))

	log.Println("Listening to clients on port 54321...")
	for {
		conn, err := listener.Accept()
		if util.CheckErr(err) {
			continue
		}
		authenticateClient(conn)

		rpc.ServeConn(conn)
	}
}
