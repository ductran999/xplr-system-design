package main

import (
	"fmt"
	"viper/interactor"
	"viper/presenter"
	"viper/router"
	"viper/view"
)

func main() {
	v := &view.ConsoleUserView{}
	i := &interactor.UserInteractor{}
	r := &router.UserRouter{}

	p := &presenter.UserPresenter{
		View:       v,
		Interactor: i,
		Router:     r,
	}

	// Scenario 1: User see Profile Page (no navigate)
	fmt.Println("[Scenario 1]: User see Profile ID 1...")
	p.OnViewProfileLoaded(1)

	fmt.Println("-------------------------------------")

	// Scenario 2: User click view detail of one user on the users list
	fmt.Println("[Scenario 2]: User click view detail User ID 1...")
	p.OnUserSelected(1)
}
