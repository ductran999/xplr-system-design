package view

import "fmt"

type UserView interface {
	ShowUserInfo(name string)
	ShowLoading()
	ShowError(message string)
}

// ConsoleUserView (Simulate Console)
type ConsoleUserView struct{}

func (v *ConsoleUserView) ShowUserInfo(name string) {
	fmt.Println("[UI DISPLAY]: User name:", name)
}

func (v *ConsoleUserView) ShowLoading() {
	fmt.Println("[UI DISPLAY]: Loading...")
}

func (v *ConsoleUserView) ShowError(message string) {
	fmt.Println("[UI DISPLAY]: Error:", message)
}
