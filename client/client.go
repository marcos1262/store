package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"store/model"
	"crypto/rsa"
	"crypto/rand"
	"store/store"
	"store/cryptopasta"
	"crypto/sha256"
	"encoding/hex"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	server, err := rpc.Dial("tcp", service)
	store.CheckMortalErr(err)

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	store.CheckMortalErr(err)
	publicKey := &key.PublicKey

	encrypted, err := cryptopasta.EncryptOAEP(publicKey, []byte("Yeah!"))
	store.CheckMortalErr(err)
	println(hex.EncodeToString(encrypted))

	decrypted, err := cryptopasta.DecryptOAEP(key, encrypted)
	store.CheckMortalErr(err)
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
	if err != nil {
		log.Fatal("Error on creating product")
	}
	println("Created", product.String())

	var res []model.Product
	queryData := model.ProductQueryData{model.Product{Idproduct: id}, 0, 1}
	err = server.Call("RPC_product.Read", queryData, &res)
	if err != nil {
		log.Fatal("Error on reading products")
	}
	if res[0].Name != "test1" {
		log.Fatal("Error on reading product name")
	}
	println("Read", res[0].String())

	var nothing int
	product = model.Product{Idproduct: id, Name: "test"}
	err = server.Call("RPC_product.Update", product, &nothing)
	if err != nil {
		log.Fatal("Error on updating product")
	}
	println("Updated", product.String())

	err = server.Call("RPC_product.Delete", product, &nothing)
	if err != nil {
		log.Fatal("Error on deleting product")
	}
	println("Deleted", product.String())
}
