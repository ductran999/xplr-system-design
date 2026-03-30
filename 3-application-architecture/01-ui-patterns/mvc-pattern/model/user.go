package model

import "errors"

type User struct {
	ID   int
	Name string
	Role string
}

func GetUserByID(id int) (*User, error) {
	if id == 1 {
		return &User{ID: 1, Name: "Alice MVC", Role: "Admin"}, nil
	}

	return nil, errors.New("user not found")
}
