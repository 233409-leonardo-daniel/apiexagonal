package repositories

type IRabbitProducer interface {
	Publish(idProduct int32, quantity int32, totalPrice float64, status string) error
}
