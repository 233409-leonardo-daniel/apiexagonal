package repositories

type IBcrypService interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword string, providedPassword string) bool
}
