package viewmodel

import (
	"mvvm/model"
	"strings"
)

type UserViewModel struct {
	FullName          string
	OnPropertyChanged func(string)
	user              model.User
}

func NewUserViewModel(u model.User) *UserViewModel {
	vm := &UserViewModel{user: u}
	vm.sync()
	return vm
}

func (vm *UserViewModel) sync() {
	vm.FullName = strings.ToUpper(vm.user.FirstName + " " + vm.user.LastName)
}

// (Simulate Two-way Binding)
func (vm *UserViewModel) SetName(first, last string) {
	vm.user.FirstName = first
	vm.user.LastName = last
	vm.sync()

	if vm.OnPropertyChanged != nil {
		vm.OnPropertyChanged(vm.FullName)
	}
}
