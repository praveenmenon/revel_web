package controllers

import "github.com/revel/revel"

type Account struct {
	*revel.Controller
}

func (c Account) Sign_up() revel.Result {
	return c.Render()
}