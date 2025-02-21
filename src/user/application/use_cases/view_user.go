package usecases

import (
	"api/src/user/domain/entities"
	"api/src/user/domain/repositories"
)

type ViewUser struct {
	db repositories.IUser
}

func NewViewUser(db repositories.IUser) *ViewUser {
	return &ViewUser{db: db}
}

func (vu *ViewUser) Execute() ([]entities.User, error) {
	return vu.db.GetAll()
}
