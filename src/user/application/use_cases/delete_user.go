package usecases

import (
	"api/src/user/domain/repositories"
)

type DeleteUser struct {
	repo repositories.IUser
}

func NewDeleteUser(repo repositories.IUser) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (uc *DeleteUser) Execute(id int32) error {
	return uc.repo.Delete(id)
}
