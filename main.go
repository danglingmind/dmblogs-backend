package main

import (
	"os"
	"strconv"

	"danglingmind.com/ddd/config"
	"danglingmind.com/ddd/infrastructure/auth"
	"danglingmind.com/ddd/infrastructure/persistence"
	"danglingmind.com/ddd/interfaces"
	"danglingmind.com/ddd/interfaces/middleware"
	"github.com/joho/godotenv"
)

func main() {

	// initialize configurations
	// conf := config.LoadConfig()
	mySqlDbname, _ := config.GetValue("mysql_db")
	mySqlUser, _ := config.GetValue("mysql_user")
	mySqlPass, _ := config.GetValue("mysql_password")
	mySqlPort, _ := config.GetValue("mysql_port")
	mySqlPortInt, _ := strconv.Atoi(mySqlPort)
	mySqlHost, _ := config.GetValue("mysql_host")
	redisHost, _ := config.GetValue("redis_host")
	redisPort, _ := config.GetValue("redis_port")
	redisPass, _ := config.GetValue("redis_password")

	// infrastructure layer instance
	// create domain repositories to deal with databases
	services, err := persistence.NewRepository(mySqlHost, mySqlUser, mySqlPass, mySqlDbname, mySqlPortInt)
	if err != nil {
		panic(err.Error())
	}

	// interface layer instance
	// create interfaces (adapters) of each interaction point to the app
	rd, err := auth.NewRedisDB(redisHost, redisPort, redisPass)
	if err != nil {
		panic(err.Error())
	}

	au := auth.NewAuth(rd.Client)
	tk := auth.NewToken()
	usersHandlers := interfaces.NewUser(services.User)
	authenticator := interfaces.NewAuthenticator(services.User, au, tk)

	// Initialize server
	server := interfaces.NewServer()
	// CORS middleware
	server.Router.Use(middleware.CORSMiddleware)

	// user service
	server.AddRoute("GET", "/users", usersHandlers.GetAllUsers)
	server.AddRoute("GET", "/user/{id:[0-9]+}", usersHandlers.GetUserById)
	server.AddRoute("POST", "/user/save", usersHandlers.Save)

	// login service
	server.AddRoute("POST", "/login", authenticator.Login)

	loginRouter := server.Router.PathPrefix("/auth/").Subrouter()
	// add authentication to login routes
	loginRouter.Use(middleware.AuthMiddleware)
	// loginRouter.HandleFunc("/login", authenticator.Login).Methods("POST")
	loginRouter.HandleFunc("/logout", authenticator.Logout).Methods("POST")
	// loginRouter.HandleFunc("/refresh", authenticator.Refresh).Methods("POST")

	// Run the server
	godotenv.Load()
	port := os.Getenv("PORT")

	server.Run(port)
}
