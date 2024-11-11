package domain

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255;unique"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	Create(user *User) error
}

type UserUsecase interface {
	GetAllUsers() ([]User, error)
	CreateUser(user *User) error
}
