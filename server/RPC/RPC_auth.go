package RPC

import (
	"log"
	"store/util"
	"errors"
)

type RPC_auth struct{}

func (r *RPC_auth) ExchangePublicKey(clientKey *[]byte, serverKey *[]byte) (err error) {
	log.Print("RPC call: Exchange public key")
	//save clientKey
	//send serverKey
	if util.CheckErr(err) {
		err = errors.New("Error on creating product: "+err.Error())
	}
	return
}
