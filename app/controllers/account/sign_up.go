package account

import "github.com/revel/revel"

type Account1 struct {
	*revel.Controller
}

func (c Account1) Sign_up() revel.Result {
	return c.Render()
}