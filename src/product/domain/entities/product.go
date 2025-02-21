package entities

type Product struct {
	Id    int32   `json:"id" gorm:"column:idProduct;primaryKey;autoIncrement"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{Name: name, Price: price}
}
