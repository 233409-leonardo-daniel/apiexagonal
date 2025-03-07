package usecases

import (
	"api/src/order/domain/repositories"
)

type DeleteOrder struct {
	repo repositories.IOrder
}

func NewDeleteOrder(repo repositories.IOrder) *DeleteOrder {
	return &DeleteOrder{repo: repo}
}

func (uc *DeleteOrder) Execute(id int32) error {
	return uc.repo.Delete(id)
}
