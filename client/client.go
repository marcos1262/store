package main

import (
	"fmt"
	"net/rpc"
	"os"
	"store/model"
	"crypto/rsa"
	"store/cryptopasta"
	"store/util"
	"net"
	"log"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	sessionKey = &[32]byte{}
)

func connect(service string) *rpc.Client {
	log.Println("Connecting to the server")

	conn, err := net.Dial("tcp", service)
	util.CheckMortalErr(err)

	authenticate(conn)

	return rpc.NewClient(conn)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	log.Println("Generating keys for cryptography")
	var err error
	privateKey, publicKey, err = cryptopasta.GenerateKeys(6144)
	util.CheckMortalErr(err)

	server := connect(service)

	var opt string
	for opt == "" {
		print("\n# PRODUCTS MANAGEMENT:\n" +
			"1 - New Product\n" +
			"2 - Show All\n" +
			"3 - Find Product\n" +
			"4 - Update Product\n" +
			"5 - Delete Product\n" +
			"6 - Exit\n" +
			"Choose an option: ")
		fmt.Scan(&opt)

	Options:
		switch opt {
		case "1":
			var id int
			product := model.Product{}

			print("Type the name: ")
			fmt.Scanln(&product.Name)

			print("Type the price: ")
			fmt.Scanln(&product.Price)

			for {
				err = server.Call("RPC_product.Create", product, &id)
				if !util.CheckErr(err) {
					break
				}
				if err == rpc.ErrShutdown {
					server = connect(service)
				} else {
					break Options
				}
			}
			product.Idproduct = id

			println("Created", product.String())

			break
		case "2":
			res := model.ProductQueryResult{}
			queryData := model.ProductQueryData{
				Product: &model.Product{},
				Start:   0, Quantity: 100,
			}

			for {
				err = server.Call("RPC_product.Read", queryData, &res)
				if !util.CheckErr(err) {
					break
				}
				if err == rpc.ErrShutdown {
					server = connect(service)
				} else {
					break Options
				}
			}

			// TODO pagination
			page := 1
			totalPages := int(res.Total / len(res.Products))
			println("\nShowing page", page, "of", totalPages)
			println(model.ProductHeader())

			for _, p := range res.Products {
				println(p.StringLine())
			}

			break
		case "3":
			res := model.ProductQueryResult{}
			product := model.Product{}

			print("Type the Id (optional): ")
			fmt.Scanln(&product.Idproduct)

			print("Type the name (optional): ")
			fmt.Scanln(&product.Name)

			print("Type the price (optional): ")
			fmt.Scanln(&product.Price)

			queryData := model.ProductQueryData{
				Product: &product,
				Start:   0, Quantity: 100,
			}

			for {
				err = server.Call("RPC_product.Read", queryData, &res)
				if !util.CheckErr(err) {
					break
				}
				if err == rpc.ErrShutdown {
					server = connect(service)
				} else {
					break Options
				}
			}

			// TODO pagination
			page := 1
			totalPages := int(res.Total / len(res.Products))
			println("\nShowing page", page, "of", totalPages)
			println(model.ProductHeader())

			for _, p := range res.Products {
				println(p.StringLine())
			}

			break
		case "4":
			var nothing int
			product := model.Product{}

			print("Type the Id: ")
			fmt.Scanln(&product.Idproduct)

			print("Type the name (optional): ")
			fmt.Scanln(&product.Name)

			print("Type the price (optional): ")
			fmt.Scanln(&product.Price)

			for {
				err = server.Call("RPC_product.Update", product, &nothing)
				if !util.CheckErr(err) {
					break
				}
				if err == rpc.ErrShutdown {
					server = connect(service)
				} else {
					break Options
				}
			}

			println("Updated", product.String())

			break
		case "5":
			var nothing int
			product := model.Product{}

			print("Type the Id: ")
			fmt.Scanln(&product.Idproduct)

			for {
				err = server.Call("RPC_product.Delete", product, &nothing)
				if !util.CheckErr(err) {
					break
				}
				if err == rpc.ErrShutdown {
					server = connect(service)
				} else {
					break Options
				}
			}

			println("Deleted", product.String())

			break
		case "6":
			println("Bye bye")
			return
			break
		default:
			println("This is not a valid option")
		}

		time.Sleep(3 * time.Second)
		opt = ""
	}
}
