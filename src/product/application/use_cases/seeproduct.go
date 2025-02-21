package usecases

import (
	"api/src/product/domain/entities"
	"api/src/product/domain/repositories"
)

type ViewProduct struct {
	db repositories.IProduct
}

func NewViewProduct(db repositories.IProduct) *ViewProduct {
	return &ViewProduct{db: db}
}

func (vp *ViewProduct) Execute() ([]entities.Product, error) {
	return vp.db.GetAll()
}
