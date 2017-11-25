package main

import (
	"testing"
	"store/server/DAO"
	"store/model"
)

func TestDAO_product(t *testing.T) {
	DAO.InitializeDB()
	defer DAO.CloseDB()

	id, err := DAO.CreateProduct(model.Product{Name: "test1", Price: 2.5})
	if err != nil {
		t.Fatal("Error on creating product")
	}

	products, err := DAO.ReadProduct(model.Product{Idproduct: id}, 0, 1)
	if err != nil {
		t.Fatal("Error on reading products")
	}
	if products[0].Name != "test1" {
		t.Fatal("Error on reading product name")
	}

	err = DAO.UpdateProduct(model.Product{Idproduct: id, Name: "test"})
	if err != nil {
		t.Fatal("Error on updating product")
	}

	products, err = DAO.ReadProduct(model.Product{Idproduct: id}, 0, 1)
	if err != nil {
		t.Fatal("Error on reading products")
	}
	if products[0].Name != "test" {
		t.Fatal("Error on updating product name")
	}

	err = DAO.DeleteProduct(model.Product{Idproduct: id})
	if err != nil {
		t.Fatal("Error on deleting product")
	}
}

func TestDAO_user(t *testing.T) {
	DAO.InitializeDB()
	defer DAO.CloseDB()

	id, err := DAO.CreateUser(model.User{Name: "test1", Login: "test", Pass: "test"})
	if err != nil {
		t.Fatal("Error on creating user")
	}

	users, err := DAO.ReadUser(model.User{Iduser: id}, 0, 1)
	if err != nil {
		t.Fatal("Error on reading users")
	}
	if users[0].Name != "test1" {
		t.Fatal("Error on reading user name")
	}

	err = DAO.UpdateUser(model.User{Iduser: id, Name: "test"})
	if err != nil {
		t.Fatal("Error on updating user")
	}

	users, err = DAO.ReadUser(model.User{Iduser: id}, 0, 1)
	if err != nil {
		t.Fatal("Error on reading users")
	}
	if users[0].Name != "test" {
		t.Fatal("Error on updating user name")
	}

	err = DAO.DeleteUser(model.User{Iduser: id})
	if err != nil {
		t.Fatal("Error on deleting user")
	}
}
