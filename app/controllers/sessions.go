package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"net/http"
	"revel_web/app/db/migrate"
)

type Session struct {
	*revel.Controller
}

func (s Session) New() revel.Result {
	migration_result := migrate.Migrate_tabels()
	fmt.Println(migration_result)
	return s.Render()
}

func (s Session) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method:", r.Method)
	// if r.Method == "POST" {
	// 	r.ParseForm()
	// 	fmt.Println("Email:", r.Form["email"])
	// 	fmt.Println("password:", r.Form["password"])
	// }
}
