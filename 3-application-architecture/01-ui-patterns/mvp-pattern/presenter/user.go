package presenter

import (
	"mvp/model"
	"mvp/view"
	"strings"
	"time"
)

type UserPresenter struct {
	view  view.UserView // Presenter holds View's Interface
	model model.User    // Presenter holds Model
}

func NewUserPresenter(v view.UserView, u model.User) *UserPresenter {
	return &UserPresenter{view: v, model: u}
}

func (p *UserPresenter) OnDisplayButtonClicked() {
	p.view.ShowLoading()
	time.Sleep(1 * time.Second) // Simulate fetching

	// Get data from model and do simple biz logic that format user name to uppercase
	displayName := strings.ToUpper(p.model.Name)

	if displayName == "" {
		p.view.ShowError("User name is Empty!")
		return
	}

	// Let view show result
	p.view.ShowUserInfo("Hello: " + displayName)
}
