package domain

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetByID(id int) (*User, error)
}

type UserService interface {
	GetDisplayName(id int) (string, error)
}
