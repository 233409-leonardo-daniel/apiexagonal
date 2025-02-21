package usecases

import (
	"api/src/user/domain/repositories"
)

type UpdateUser struct {
	repo repositories.IUser
}

func NewUpdateUser(repo repositories.IUser) *UpdateUser {
	return &UpdateUser{repo: repo}
}

func (uc *UpdateUser) Execute(id int32, name string, lastname string, age int) error {
	return uc.repo.Update(id, name, lastname, age)
}
