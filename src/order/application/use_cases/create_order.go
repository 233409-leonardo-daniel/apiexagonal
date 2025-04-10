package usecases

import (
	"api/src/order/domain/repositories"
)

type IOrder interface {
	Execute(idProduct int32, quantity int32, totalPrice float64, status string) error
}

type CreateOrder struct {
	db             repositories.IOrder
	rabbitProducer repositories.IRabbitProducer
}

func NewCreateOrder(db repositories.IOrder, rabbitProducer repositories.IRabbitProducer) *CreateOrder {
	return &CreateOrder{db: db, rabbitProducer: rabbitProducer}
}

func (co *CreateOrder) Execute(idProduct int32, quantity int32, totalPrice float64, status string) error {
	// Publicar en RabbitMQ
	err := co.rabbitProducer.Publish(idProduct, quantity, totalPrice, status)
	if err != nil {
		return err
	}

	// Guardar en la base de datos
	return co.db.Save(idProduct, quantity, totalPrice, status)
}
