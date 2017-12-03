package util

import (
	"log"
	"os"
	"bufio"
	"encoding/json"
)

const (
	ERR_NOT_AUTH = "login or password incorrect"
)

func CheckErr(err error) bool {
	if err != nil {
		log.Println("Error: " + err.Error())
		return true
	}
	return false
}

func CheckMortalErr(err error) {
	if CheckErr(err) {
		os.Exit(1)
	}
}

// Receive and process a certain amount of data from in
func ReceiveData(in *bufio.Reader, qtd int) ([]byte, error) {
	data := make([]byte, qtd+1) // in buffer
	n, err := in.Read(data)     // Reading from network
	if err != nil {
		return nil, err
	}
	return data[:n-1], err
}

// Receive and process data until line-break from in
func ReceiveDataLine(in *bufio.Reader) ([]byte, error) {
	data, err := in.ReadBytes('\n') // Reading from network
	if err != nil {
		return nil, err
	}
	return data[:len(data)-1], err
}

// Prepare and send data to out
func SendData(out *bufio.Writer, data []byte) (error) {
	_, err := out.Write(append(data, '\n'))
	out.Flush()
	return err
}

func StructToBytes(s interface{}) ([]byte, error) {
	return json.Marshal(s)
}

func BytesToStruct(data []byte, s interface{}) error {
	return json.Unmarshal(data, s)
}

//func SendEncrypted(out *bufio.Writer, publicKey *rsa.PublicKey, data interface{}) error {
//
//}
