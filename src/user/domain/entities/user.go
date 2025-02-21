package entities

type User struct {
	ID       int32  `json:"id" gorm:"column:id_user;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	LastName string `json:"lastname" gorm:"column:lastname"`
	Age      int    `json:"age"`
}

func NewUser(name string, lastname string, age int) *User {
	return &User{Name: name, LastName: lastname, Age: age}
}
