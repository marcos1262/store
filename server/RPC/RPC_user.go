package RPC

import (
	"store/model"
	"store/server/DAO"
	"log"
)

type RPC_user struct{}

func (r *RPC_user) Create(p *model.User, id *int) (err error) {
	log.Print("RPC call: Create user")
	*id, err = DAO.CreateUser(p)
	return
}

func (r *RPC_user) Read(q *model.UserQueryData, result *model.UserQueryResult) (err error) {
	log.Print("RPC call: Read user")
	users, total, err := DAO.ReadUser(q.User, q.Start, q.Quantity)
	result.Users = users
	result.Total = total
	return
}

func (r *RPC_user) Update(p *model.User, nothing *int) (err error) {
	log.Print("RPC call: Update user")
	err = DAO.UpdateUser(p)
	return
}

func (r *RPC_user) Delete(p *model.User, nothing *int) (err error) {
	log.Print("RPC call: Delete user")
	err = DAO.DeleteUser(p)
	return
}
