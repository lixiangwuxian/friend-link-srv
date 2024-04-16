package main

import (
	"log"

	db "lxtend.com/friend_link/DB"
	configParser "lxtend.com/friend_link/config"
	emailhandler "lxtend.com/friend_link/email_handler"
	"lxtend.com/friend_link/webserver"
)

var sqlite3Path = "./data.db"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	emailhandler.Init(configParser.GetConfig().Email.Server,
		configParser.GetConfig().Email.Port,
		configParser.GetConfig().Email.User,
		configParser.GetConfig().Email.Password,
	)
	if err := db.InitDB(sqlite3Path); err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	webserver.StartUp()
}
