package model

import (
	"strconv"
	"fmt"
	"net"
	"crypto/rsa"
	"bufio"
)

// Product on database
type Product struct {
	Idproduct int
	Name      string
	Price     float32
}

func (p Product) String() string {
	return "Product [" +
		"\n\t" + "id: " + strconv.Itoa(p.Idproduct) +
		"\n\t" + "name: " + p.Name +
		"\n\t" + "price: " + fmt.Sprint(p.Price) +
		"\n]"
}

func (p Product) StringLine() string {
	return "Product [" +
		"\t" + strconv.Itoa(p.Idproduct) +
		"\t" + p.Name +
		"\t" + fmt.Sprint(p.Price) +
		"\t]"
}

func ProductHeader() string {
	return "         \tId\tName\tPrice"
}

// User on database
type User struct {
	Iduser int
	Name   string
	Login  string
	Pass   string
}

func (u User) String() string {
	return "User [" +
		"\n\t" + "id: " + strconv.Itoa(u.Iduser) +
		"\n\t" + "name: " + u.Name +
		"\n\t" + "login: " + u.Login +
		"\n\t" + "pass: " + u.Pass +
		"\n]"
}

func (u User) StringLine() string {
	return "User [" +
		"\t" + strconv.Itoa(u.Iduser) +
		"\t" + u.Name +
		"\t" + u.Login +
		"\t" + u.Pass +
		"\t]"
}

func UserHeader() string {
	return "      \tId\tName\t\t\tLogin\t\t\tPass"
}

// Client connected to RPC
type Client struct {
	Conn       net.Conn
	PublicKey  *rsa.PublicKey
	SessionKey *[32]byte
	In         *bufio.Reader
	Out        *bufio.Writer
}

// RPC Auxiliar type
type ProductQueryData struct {
	Product  *Product
	Start    int
	Quantity int
}

// RPC Auxiliar type
type ProductQueryResult struct {
	Products []Product
	Total    int
}

// RPC Auxiliar type
type UserQueryData struct {
	User     *User
	Start    int
	Quantity int
}

// RPC Auxiliar type
type UserQueryResult struct {
	Users []User
	Total int
}
