package usecases

import (
	"api/src/product/domain/repositories"
)

type DeleteProduct struct {
	repo repositories.IProduct
}

func NewDeleteProduct(repo repositories.IProduct) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (uc *DeleteProduct) Execute(id int32) error {
	return uc.repo.Delete(id)
}
