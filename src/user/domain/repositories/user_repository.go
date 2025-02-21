package repositories

import "api/src/user/domain/entities"

type IUser interface {
	Save(name string, lastName string, age int) error
	GetAll() ([]entities.User, error)
	Update(id int32, name string, lastname string, age int) error
	Delete(id int32) error
}
