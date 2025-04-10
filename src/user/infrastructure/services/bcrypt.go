package services

import (
	"api/src/user/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	cost int
}

func NewBcryptService() repositories.IBcrypService {
	return &BcryptService{cost: 12}
}

func (b *BcryptService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *BcryptService) ComparePasswords(hashedPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}
