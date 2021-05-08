package interfaces

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// server
type Server struct {
	Router *mux.Router
}

type baseHandler func(http.ResponseWriter, *http.Request)

func NewServer() *Server {
	return &Server{
		Router: mux.NewRouter(),
	}
}

func (s *Server) AddRoute(method, route string, handler baseHandler, args ...interface{}) {
	if len(route) > 0 && handler != nil {
		method = strings.ToUpper(method)
		switch method {
		case "GET", "POST", "PUT", "DELETE":
			s.Router.HandleFunc(route, handler).Methods(method)
		default:
			s.Router.HandleFunc(route, handler).Methods("GET")
		}
	}
}

func (s *Server) Run(port string) {
	fmt.Printf("\nStarting server at :%s\n", port)
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", "0.0.0.0", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.Router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// take CTL+C to stop the server after a 2 minutes
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// wait for 2 minutes to stop the server
	var wait time.Duration = 2 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// graceful shutdown the server
	srv.Shutdown(ctx)
	log.Println("shutting down the server")
	os.Exit(0)
}
