package router

import "fmt"

type UserRouter struct{}

func (r *UserRouter) NavigateToDetails(userName string) {
	fmt.Println("[ROUTER]: Redirect to " + userName + " profile page:...")
}
