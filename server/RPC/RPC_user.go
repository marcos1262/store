package RPC

import (
	"store/model"
	"store/server/DAO"
	"log"
)

type RPC_user struct{}

type UserQueryData struct {
	user model.User
	start int
	quantity int
}

func (r *RPC_user) Create(p *model.User, id *int) (err error) {
	log.Print("RPC call: Create user")
	*id, err = DAO.CreateUser(*p)
	return
}

func (r *RPC_user) Read(q *UserQueryData, users *[]model.User) (err error) {
	log.Print("RPC call: Read user")
	*users, err = DAO.ReadUser(q.user, q.start, q.quantity)
	return
}

func (r *RPC_user) Update(p *model.User, nothing *int) (err error) {
	log.Print("RPC call: Update user")
	err = DAO.UpdateUser(*p)
	return
}

func (r *RPC_user) Delete(p *model.User, nothing *int) (err error) {
	log.Print("RPC call: Delete user")
	err = DAO.DeleteUser(*p)
	return
}