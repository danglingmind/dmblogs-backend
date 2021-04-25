package main

import (
	"strconv"

	"danglingmind.com/ddd/v1/config"
	"danglingmind.com/ddd/v1/infrastructure/persistence"
	"danglingmind.com/ddd/v1/interfaces"
)

func main() {

	// initialize configurations
	// conf := config.LoadConfig()
	mySqlDbname, _ := config.GetValue("MysqlDbName")
	mySqlUser, _ := config.GetValue("MysqlUser")
	mySqlPass, _ := config.GetValue("Mysqlpassword")
	mySqlPort, _ := config.GetValue("MysqlPort")
	mySqlPortInt, _ := strconv.Atoi(mySqlPort)
	mySqlHost, _ := config.GetValue("MysqlHost")

	// infrastructure layer instance
	// create domain repositories to deal with databases
	services, err := persistence.NewRepository(mySqlHost, mySqlUser, mySqlPass, mySqlDbname, mySqlPortInt)
	if err != nil {
		panic(err.Error())
	}

	// interface layer instance
	// create interfaces (adapters) of each interaction point to the app
	usersHandlers := interfaces.NewUser(services.User)
	// initialize paths
	server := interfaces.NewServer()
	server.AddRoute("GET", "/users", usersHandlers.GetAllUsers) // TODO: move this function conversion into Server.AddRoute
	server.AddRoute("GET", "/user/:id", usersHandlers.GetUserById)
	server.AddRoute("POST", "/user/save", usersHandlers.Save)

	// Run the server
	server.Run(8000)
}
