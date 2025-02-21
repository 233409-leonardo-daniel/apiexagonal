package repositories

import "api/src/product/domain/entities"

type IProduct interface {
	Save(name string, price float32) error
	GetAll() ([]entities.Product, error)
	Update(id int32, name string, price float32) error
	Delete(id int32) error
}
