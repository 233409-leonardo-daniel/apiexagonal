package usecases

import "api/src/user/domain/repositories"

type CreateUser struct {
	db repositories.IUser
}

func NewCreateUser(db repositories.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cp *CreateUser) Execute(name string, lastname string, age int) error {
	return cp.db.Save(name, lastname, age)
}
