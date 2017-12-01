package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"store/model"
	"crypto/rsa"
	"crypto/rand"
	"store/util"
	"store/cryptopasta"
	"crypto/sha256"
	"encoding/hex"
	"net"
	"bufio"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	util.CheckMortalErr(err)
	publicKey := &key.PublicKey

	conn, err := net.Dial("tcp", service)
	util.CheckMortalErr(err)

	in := bufio.NewReader(conn)
	out := bufio.NewWriter(conn)

	// Send public key, receive server's public key
	var serverKey []byte
	product := model.Product{Name: "test1", Price: 2.5}
	err = server.Call("RPC_auth.ExchangePublicKey", publicKey, &serverKey)
	util.CheckErr(err)
	println("Created", product.String())

	// Send auth info encrypted, receive session key or error (if not authenticate)



	server, err := rpc.Dial("tcp", service)
	util.CheckMortalErr(err)



	encrypted, err := cryptopasta.EncryptOAEP(publicKey, []byte("Yeah!"))
	util.CheckMortalErr(err)
	println(hex.EncodeToString(encrypted))

	decrypted, err := cryptopasta.DecryptOAEP(key, encrypted)
	util.CheckMortalErr(err)
	println(string(decrypted))

	sessionKey := sha256.Sum256([]byte("password"))
	println(hex.EncodeToString(sessionKey[:]))

	encrypted, err = cryptopasta.EncryptAES([]byte("Yeah!"), &sessionKey)
	println(hex.EncodeToString(encrypted))

	decrypted, err = cryptopasta.DecryptAES(encrypted, &sessionKey)
	println(string(decrypted))

	os.Exit(0)
	//////////////////////////////////////////////////////////////////////////////////

	var id int
	product := model.Product{Name: "test1", Price: 2.5}
	err = server.Call("RPC_product.Create", product, &id)
	util.CheckErr(err)
	println("Created", product.String())

	var res []model.Product
	queryData := model.ProductQueryData{model.Product{Idproduct: id}, 0, 1}
	err = server.Call("RPC_product.Read", queryData, &res)
	util.CheckErr(err)
	if res[0].Name != "test1" {
		log.Fatal("Error on reading product name")
	}
	println("Read", res[0].String())

	var nothing int
	product = model.Product{Idproduct: id, Name: "test"}
	err = server.Call("RPC_product.Update", product, &nothing)
	util.CheckErr(err)
	println("Updated", product.String())

	err = server.Call("RPC_product.Delete", product, &nothing)
	util.CheckErr(err)
	println("Deleted", product.String())
}
