package main

import (
	"os"

	"danglingmind.com/ddd/infrastructure/auth"
	"danglingmind.com/ddd/infrastructure/persistence"
	"danglingmind.com/ddd/interfaces"
	"danglingmind.com/ddd/interfaces/middleware"
	"github.com/joho/godotenv"
)

func main() {

	// initialize configurations
	godotenv.Load()
	// db config
	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	// redis config
	redis_url := os.Getenv("REDISTOGO_URL")
	// redis_host := os.Getenv("REDIS_HOST")
	// redis_port := os.Getenv("REDIS_PORT")
	// redis_password := os.Getenv("REDIS_PASSWORD")


	// infrastructure layer instance
	// create domain repositories to deal with databases
	services, err := persistence.NewRepositories(dbdriver, user, password, dbport, host, dbname)
	if err != nil {
		panic(err.Error())
	}

	// interface layer instance
	// create interfaces (adapters) of each interaction point to the app
	rd, err := auth.NewRedisDB(redis_url)

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
	port := os.Getenv("PORT")

	server.Run(port)
}
