package main

import (
	"store/model"
)

var clients = make(map[string]*model.Client)

type keyValue struct {
	key    string
	client *model.Client
}

var (
	removeClient chan string
	findClient   chan keyValue
	newClient    chan keyValue
)

func manageClients() {
	removeClient = make(chan string)
	findClient = make(chan keyValue)
	newClient = make(chan keyValue)

	for {
		select {
		// Handle Clients
		case kv := <-newClient:
			clients[kv.key] = kv.client

		case key := <-removeClient:
			delete(clients, key)

		case kv := <-findClient:
			findClient <- keyValue{kv.key, clients[kv.key]}
		}
	}
}

func saveClient(key string, client *model.Client) {
	newClient <- keyValue{key, client}
}

func getClient(key string) *model.Client {
	findClient <- keyValue{key: key}
	return (<-findClient).client
}

func delClient(key string) {
	removeClient <- key
}
