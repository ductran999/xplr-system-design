package view

import "fmt"

type UserView interface {
	DisplayUser(name string)
}

type ConsoleUserView struct{}

func (v *ConsoleUserView) DisplayUser(name string) {
	fmt.Println("[VIEW]: User:", name)
}
