package usecases

import "api/src/product/domain/repositories"

type CreateProduct struct {
	db repositories.IProduct
}

func NewCreateProduct(db repositories.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (cp *CreateProduct) Execute(name string, price float32) error {
	return cp.db.Save(name, price)
}
