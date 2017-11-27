package RPC

import (
	"store/model"
	"store/server/DAO"
	"log"
)

type RPC_product struct{}

func (r *RPC_product) Create(p *model.Product, id *int) (err error) {
	log.Print("RPC call: Create product")
	*id, err = DAO.CreateProduct(*p)
	return
}

func (r *RPC_product) Read(q *model.ProductQueryData, products *[]model.Product) (err error) {
	log.Print("RPC call: Read product")
	*products, err = DAO.ReadProduct(q.Product, q.Start, q.Quantity)
	return
}

func (r *RPC_product) Update(p *model.Product, nothing *int) (err error) {
	log.Print("RPC call: Update product")
	err = DAO.UpdateProduct(*p)
	return
}

func (r *RPC_product) Delete(p *model.Product, nothing *int) (err error) {
	log.Print("RPC call: Delete product")
	err = DAO.DeleteProduct(*p)
	return
}
