package main

import (
	"fmt"
	"mvvm/model"
	"mvvm/view"
	"mvvm/viewmodel"
)

func main() {
	u := model.User{FirstName: "Alice", LastName: "Smith"}

	vm := viewmodel.NewUserViewModel(u)

	vm.OnPropertyChanged = func(newName string) {
		view.Render(newName)
	}

	view.Render(vm.FullName)

	fmt.Println("\n(Action): User change name to 'Danny Ng'...")
	vm.SetName("Danny", "Ng") // the view will re-render with new value
}
