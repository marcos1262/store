package RPC

import (
	"store/model"
	"store/server/DAO"
)

type RPC_product struct{}

type ProductQueryData struct {
	product model.Product
	start int
	quantity int
}

func (r *RPC_product) Create(p *model.Product, id *int) (err error) {
	*id, err = DAO.CreateProduct(*p)
	return
}

func (r *RPC_product) Read(q *ProductQueryData, products *[]model.Product) (err error) {
	*products, err = DAO.ReadProduct(q.product, q.start, q.quantity)
	return
}

func (r *RPC_product) Update(p *model.Product, nothing *int) (err error) {
	err = DAO.UpdateProduct(*p)
	return
}

func (r *RPC_product) Delete(p *model.Product, nothing *int) (err error) {
	err = DAO.DeleteProduct(*p)
	return
}