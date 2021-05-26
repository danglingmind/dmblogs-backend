package main

import (
	"os"

	"danglingmind.com/ddd/domain/service"
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
	redisHost := os.Getenv("REDIS_HOST")

	// create domain repositories to deal with databases
	services, err := persistence.NewRepositories(dbdriver, user, password, dbport, host, dbname)
	if err != nil {
		panic(err.Error())
	}

	// create interfaces (adapters) of each interaction point to the app
	rd, err := auth.NewRedisDB(redisHost)
	if err != nil {
		panic(err.Error())
	}

	au := auth.NewAuth(rd.Client)
	tk := auth.NewToken()
	buServices := service.NewTagService(services.Tag, services.BlogTag)
	usersHandlers := interfaces.NewUser(services.User)
	authenticator := interfaces.NewAuthenticator(services.User, au, tk)
	blogHandlers := interfaces.NewBlog(
		services.Blog,
		services.User,
		services.Tag,
		services.BlogTag,
		buServices,
		tk,
		au,
	)

	// Initialize server
	server := interfaces.NewServer()
	// CORS middleware
	server.Router.Use(middleware.CORSMiddleware)

	// login service endpoints
	server.AddRoute("PUT", "/register", authenticator.Register)
	server.AddRoute("POST", "/login", authenticator.Login)

	// add authentication to login routes
	authenticatedRouter := server.Router.PathPrefix("/auth/").Subrouter()
	authenticatedRouter.Use(middleware.AuthMiddleware)

	authenticatedRouter.HandleFunc("/logout", authenticator.Logout).Methods("POST")
	// authenticatedRouter.HandleFunc("/refresh", authenticator.Refresh).Methods("POST")

	// user service endpoints
	server.AddRoute("GET", "/users", usersHandlers.GetAllUsers)
	server.AddRoute("GET", "/users/{id:[0-9]+}", usersHandlers.GetUserById)

	// blogs endpoints
	authenticatedRouter.
		Path("/blogs/save").
		Methods("PUT").
		HandlerFunc(blogHandlers.Save).
		Name("SaveBlog")

	authenticatedRouter.
		Path("/blogs/{id:[0-9]+}").
		Methods("GET").
		HandlerFunc(blogHandlers.GetBlogById).
		Name("GetBlogById")

	// support limit and offset query params
	authenticatedRouter.
		Path("/blogs/tag/{tagid:[0-9]+}").
		Methods("GET").
		HandlerFunc(blogHandlers.GetBlogsByTagName).
		Name("BlogsByTag")

	// support limit and offset query params
	authenticatedRouter.
		Path("/blogs").
		Methods("GET").
		HandlerFunc(blogHandlers.GetBlogs).
		Name("BlogsPage")

	// tag service endpoints

	// Run the server
	port := os.Getenv("PORT")

	server.Run(port)
}
