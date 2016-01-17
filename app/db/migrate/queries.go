package migrate

import (
	_ "github.com/lib/pq"
	"log"
	"revel_web/app"
)

func Migrate_tabels() string {

	var result string

	app.InitDB()

	if _, err := app.DB.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL, firstname varchar(100), lastname varchar(100), email varchar(320), user_thumbnail varchar(2083), password varchar(100), password_confirmation varchar(100), city varchar(100), state varchar(100), country varchar(100), createdat time, updatedat time , description varchar(320) DEFAULT 'Im rViDi user', devise_token varchar(320), PRIMARY KEY(id))"); err != nil {
		log.Fatal(err)
	} else {
		result = "Users table created. "
	}
	devices, err := app.DB.Exec("CREATE TABLE IF NOT EXISTS devices (id SERIAL, devise_token varchar(320), user_id int, PRIMARY KEY(devise_token), CONSTRAINT uid_key FOREIGN KEY(user_id) REFERENCES users(id))")
	if err != nil || devices == nil {
		log.Fatal(err)
	} else {
		result = result + "Devices table created. "
	}
	sessions, err := app.DB.Exec("CREATE TABLE IF NOT EXISTS sessions (id int,start_time timestamptz,end_time timestamptz,user_id int,CONSTRAINT session_id_key FOREIGN KEY (user_id) REFERENCES users (id), devise_token varchar(320), CONSTRAINT dev_id_key FOREIGN KEY(devise_token) REFERENCES devices (devise_token))")
	if err != nil || sessions == nil {
		log.Fatal(err)
	} else {
		result = result + "Sessions table created. "
	}
	shows, err := app.DB.Exec("CREATE TABLE IF NOT EXISTS shows (id SERIAL,title varchar(255),tag varchar(255),length bigint,cameo_length bigint DEFAULT 0,display_privacy varchar(255) NOT NULL CHECK (display_privacy IN ('private','public')), contribution_privacy varchar(255) NOT NULL CHECK (contribution_privacy IN ('private','public')),description varchar(255),review_cameo boolean DEFAULT FALSE,show_url varchar(2083),show_thumbnail varchar(2083),user_id int,status boolean DEFAULT TRUE,PRIMARY KEY(id),CONSTRAINT uid_shows_key FOREIGN KEY (user_id)REFERENCES users(id))")
	if err != nil || shows == nil {
		log.Fatal(err)
	} else {
		result = result + "Shows table created. "
	}
	invitees, err := app.DB.Exec("CREATE TABLE IF NOT EXISTS invitees (id SERIAL, sender_user_id int, reciever_user_id int, invitation_type varchar(255), show_id int, status boolean DEFAULT FALSE, CONSTRAINT sender_id_key FOREIGN KEY(sender_user_id) REFERENCES users(id), CONSTRAINT reciever_id_key FOREIGN KEY(reciever_user_id) REFERENCES users(id), PRIMARY KEY(id), CONSTRAINT show_id_key FOREIGN KEY (show_id) REFERENCES shows (id))")
	if err != nil || invitees == nil {
		log.Fatal(err)
	} else {
		result = result + "Invitees table created. "
	}
	notificationss, err := app.DB.Exec("CREATE TABLE IF NOT EXISTS notifications (id SERIAL,invitation_id int,created_at timestamptz,status boolean DEFAULT TRUE,PRIMARY KEY(id),CONSTRAINT invitation_id_notifications FOREIGN KEY (invitation_id) REFERENCES invitees(id))")
	if err != nil || notificationss == nil {
		log.Fatal(err)
	} else {
		result = result + "Notifications table created. "
	}

	result = result + "Migrations are up to date."

	return result
}
