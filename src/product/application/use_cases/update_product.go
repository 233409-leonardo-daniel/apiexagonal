package usecases

import (
	"api/src/product/domain/repositories"
)

type UpdateProduct struct {
	repo repositories.IProduct
}

func NewUpdateProduct(repo repositories.IProduct) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (uc *UpdateProduct) Execute(id int32, name string, price float32) error {
	return uc.repo.Update(id, name, price)
}
