package usecases

import (
	"api/src/user/domain/repositories"
)

type CreateUser struct {
	db           repositories.IUser
	cryptService repositories.IBcrypService
}

func NewCreateUser(db repositories.IUser, cryptService repositories.IBcrypService) *CreateUser {
	return &CreateUser{
		db:           db,
		cryptService: cryptService,
	}
}

func (cp *CreateUser) Execute(name string, lastname string, age int, password string) error {
	hashedPassword, err := cp.cryptService.HashPassword(password)
	if err != nil {
		return err
	}
	return cp.db.Save(name, lastname, age, hashedPassword)
}
