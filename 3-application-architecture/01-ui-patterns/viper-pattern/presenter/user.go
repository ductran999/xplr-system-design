package presenter

import (
	"viper/interactor"
	"viper/router"
	"viper/view"
)

type UserPresenter struct {
	View       view.UserView
	Interactor *interactor.UserInteractor
	Router     *router.UserRouter
}

func (p *UserPresenter) OnViewProfileLoaded(id int) {
	user, err := p.Interactor.FetchUserByID(id)
	if err != nil {
		return
	}

	p.View.DisplayUser(user.Name)
}

func (p *UserPresenter) OnUserSelected(id int) {
	user, err := p.Interactor.FetchUserByID(id)
	if err != nil {
		return
	}

	p.Router.NavigateToDetails(user.Name)
}
