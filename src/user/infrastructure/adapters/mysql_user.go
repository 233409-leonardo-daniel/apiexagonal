package adapters

import (
	"api/src/user/domain/entities"
	"api/src/user/domain/repositories"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) repositories.IUser {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Save(name string, lastname string, age int, password string) error {
	user := entities.NewUser(name, lastname, age, password)
	result := repo.db.Create(user)
	return result.Error
}

func (repo *MySQLRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	result := repo.db.Find(&users)
	return users, result.Error
}

func (repo *MySQLRepository) Update(id int32, name string, lastname string, age int) error {
	var user entities.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	user.Name = name
	user.LastName = lastname
	user.Age = age

	result = repo.db.Save(&user)
	return result.Error
}

func (repo *MySQLRepository) Delete(id int32) error {
	var user entities.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.db.Delete(&user)
	return result.Error
}
