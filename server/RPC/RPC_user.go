package RPC

import (
	"store/model"
	"store/server/DAO"
)

type RPC_user struct{}

type UserQueryData struct {
	user model.User
	start int
	quantity int
}

func (r *RPC_user) Create(p *model.User, id *int) (err error) {
	*id, err = DAO.CreateUser(*p)
	return
}

func (r *RPC_user) Read(q *UserQueryData, users *[]model.User) (err error) {
	*users, err = DAO.ReadUser(q.user, q.start, q.quantity)
	return
}

func (r *RPC_user) Update(p *model.User, nothing *int) (err error) {
	err = DAO.UpdateUser(*p)
	return
}

func (r *RPC_user) Delete(p *model.User, nothing *int) (err error) {
	err = DAO.DeleteUser(*p)
	return
}