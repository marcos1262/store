package RPC

import (
	"store/model"
	"store/server/DAO"
	"log"
	"store/util"
	"errors"
)

type RPC_product struct{}

func (r *RPC_product) Create(p *model.Product, id *int) (err error) {
	log.Print("RPC call: Create product")
	*id, err = DAO.CreateProduct(p)
	if util.CheckErr(err) {
		err = errors.New("Error on creating product: " + err.Error())
	}
	return
}

func (r *RPC_product) Read(q *model.ProductQueryData, result *model.ProductQueryResult) (err error) {
	log.Print("RPC call: Read product")
	products, total, err := DAO.ReadProduct(q.Product, q.Start, q.Quantity)
	if util.CheckErr(err) {
		err = errors.New("Error on reading products: " + err.Error())
	}
	result.Products = products
	result.Total = total
	return
}

func (r *RPC_product) Update(p *model.Product, nothing *int) (err error) {
	log.Print("RPC call: Update product")
	err = DAO.UpdateProduct(p)
	if util.CheckErr(err) {
		err = errors.New("Error on updating product: " + err.Error())
	}
	return
}

func (r *RPC_product) Delete(p *model.Product, nothing *int) (err error) {
	log.Print("RPC call: Delete product")
	err = DAO.DeleteProduct(p)
	if util.CheckErr(err) {
		err = errors.New("Error on deleting product: " + err.Error())
	}
	return
}
