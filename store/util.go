package store

import (
	"log"
	"os"
)

func CheckErr(err error) bool {
	if err != nil {
		log.Fatal("Error: "+err.Error())
		return true
	}
	return false
}

func CheckMortalErr(err error) {
	if CheckErr(err) {
		os.Exit(1)
	}
}