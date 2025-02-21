package adapters

import (
	"api/src/product/domain/entities"
	"api/src/product/domain/repositories"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) repositories.IProduct {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Save(name string, price float32) error {
	product := entities.NewProduct(name, price)
	result := repo.db.Create(product)
	return result.Error
}

func (repo *MySQLRepository) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	result := repo.db.Find(&products)
	return products, result.Error
}

func (repo *MySQLRepository) Update(id int32, name string, price float32) error {
	var product entities.Product
	result := repo.db.First(&product, id)
	if result.Error != nil {
		return result.Error
	}

	product.Name = name
	product.Price = price

	result = repo.db.Save(&product)
	return result.Error
}

func (repo *MySQLRepository) Delete(id int32) error {
	var product entities.Product
	result := repo.db.First(&product, id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.db.Delete(&product)
	return result.Error
}
