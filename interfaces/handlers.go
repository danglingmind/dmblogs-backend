package interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// printDebugf behaves like log.Printf only in the debug env
func printDebugf(format string, args ...interface{}) {
	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// ErrorResponse is Error response template
type ErrorResponse struct {
	Message string `json:"reason"`
	Error   error  `json:"-"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("reason: %s, error: %s", e.Message, e.Error.Error())
}

// Respond is response write to ResponseWriter
func Respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, err, "invalid json")
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *ErrorResponse, ErrorResponse:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, err, "failed to parse json")
			return
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int, err error, msg string) {
	e := &ErrorResponse{
		Message: msg,
		Error:   err,
	}
	printDebugf("%s", e.String())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, e)
}

// JSON is wrapped Respond when success response
func JSON(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, src)
}

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
