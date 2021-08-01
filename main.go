package main

import (
	"log"
	"os"

	"danglingmind.com/ddd/domain/service"
	"danglingmind.com/ddd/infrastructure/auth"
	"danglingmind.com/ddd/infrastructure/persistence"
	"danglingmind.com/ddd/interfaces"
	"danglingmind.com/ddd/interfaces/middleware"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	// initialize configurations
	godotenv.Load()
	// initialize the log
	logInstance := logrus.New()
	logInstance.SetFormatter(&logrus.JSONFormatter{})

	log.SetOutput(logInstance.Writer())
	logrus.SetOutput(logInstance.Writer())
	// pass our global logger to the middleware as well
	logMiddleware := middleware.LoggingMiddleware(logInstance)

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

	// add api prefix to all api endpoints
	apiSubrouter := server.Router.PathPrefix("/api/").Subrouter()
	// CORS middleware
	apiSubrouter.Use(middleware.CORSMiddleware)
	// logging middleware
	apiSubrouter.Use(logMiddleware)

	// login service endpoints
	apiSubrouter.
		Path("/users/register").
		Methods("PUT").
		HandlerFunc(authenticator.Register)

	apiSubrouter.
		Path("/users/login").
		Methods("POST").
		HandlerFunc(authenticator.Login)

	// add authentication to login routes
	authenticatedRouter := apiSubrouter.PathPrefix("/auth/").Subrouter()
	authenticatedRouter.Use(middleware.AuthMiddleware)

	authenticatedRouter.
		Path("/logout").
		Methods("POST").
		HandlerFunc(authenticator.Logout)

	authenticatedRouter.
		Path("/refresh").
		Methods("POST").
		HandlerFunc(authenticator.Refresh)

	// user service endpoints
	apiSubrouter.
		Path("/users").
		Methods("GET").
		HandlerFunc(usersHandlers.GetAllUsers)

	apiSubrouter.
		Path("/users/{id}").
		Methods("GET").
		HandlerFunc(usersHandlers.GetUserById)

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

	// get user's blog
	authenticatedRouter.
		Path("/blogs/user/{id:[0-9]+}").
		Methods("GET").
		HandlerFunc(blogHandlers.GetBlogByUserId).
		Name("GetBlogsByUserId")

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
