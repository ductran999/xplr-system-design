package main

import (
	"fmt"
	"mvp/model"
	"mvp/presenter"
	"mvp/view"
)

func main() {
	user := model.User{ID: 1, Name: "Alice MVP"}

	ui := &view.ConsoleUserView{}

	p := presenter.NewUserPresenter(ui, user)

	fmt.Println("(Action): User click 'User profile' button...")
	p.OnDisplayButtonClicked()
}
