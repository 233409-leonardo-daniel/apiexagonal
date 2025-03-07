package usecases

import (
	"api/src/order/domain/entities"
	"api/src/order/domain/repositories"
)

type ViewOrder struct {
	db repositories.IOrder
}

func NewViewOrder(db repositories.IOrder) *ViewOrder {
	return &ViewOrder{db: db}
}

func (vo *ViewOrder) Execute() ([]entities.Order, error) {
	return vo.db.GetAll()
}
