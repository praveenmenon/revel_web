package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"log"
	"revel_web/app"
	"revel_web/app/db/migrate"
)

type Account struct {
	*revel.Controller
}

func (c Account) Sign_up() revel.Result {
	migration_result := migrate.Migrate_tabels()
	fmt.Println(migration_result)

	session_response, err := app.DB.Query("SELECT devise_token,user_id from sessions")
	if err != nil {
		log.Fatal(err)
	}
	for session_response.Next() {
		var devise_token string
		var id int
		err := session_response.Scan(&devise_token, &id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("devise_token:", devise_token)
		fmt.Println("id:", id)
	}
	return c.Render()
}
