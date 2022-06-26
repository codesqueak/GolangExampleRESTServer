package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restServer/handlers"
	"restServer/model"
	"syscall"
	"time"
)

var webServer *http.Server

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		model.Port = ":" + arguments[1]
	} else {
		model.Port = ":8090"
	}
	//
	// Create the web server
	router := mux.NewRouter()
	webServer = &http.Server{
		Addr:         model.Port,
		Handler:      router,
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	// add the stuff we want to serve
	addRoutes(router)
	addSwaggerSupport(router)
	// start
	go server()
	// ... and give us a way to stop it
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // os signals
	log.Print("Server Started on port", model.Port)
	<-shutdown
	log.Print("Server Stopped")
}

// RESTful Endpoints
func addRoutes(router *mux.Router) {
	router.HandleFunc("/adduser", handlers.AddUserPostHandler).Methods(http.MethodPost)                // keep adding users
	router.HandleFunc("/adduser/{uuid}", handlers.AddUserPutHandler).Methods(http.MethodPut)           // add a specific user
	router.HandleFunc("/getuser", handlers.GetUserWithQueryValueHandler).Methods(http.MethodGet)       // get a user
	router.HandleFunc("/getuser/{uuid}", handlers.GetUserArgInURLHandler).Methods(http.MethodGet)      // get a user
	router.HandleFunc("/getusers", handlers.GetAllUsersHandler).Methods(http.MethodGet)                // get a user
	router.HandleFunc("/getusers", handlers.GetAllUsersHeadHandler).Methods(http.MethodHead)           // get a user
	router.HandleFunc("/updateuser/{uuid}", handlers.UpdateUserPatchHandler).Methods(http.MethodPatch) // update user data
	// fallback
	router.HandleFunc("/", handlers.DefaultHandler)
}

// Swagger support
func addSwaggerSupport(router *mux.Router) {
	//
	// Server the swagger.yaml file
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	uiOpts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	//
	// Set up the swagger UI
	swaggerUI := middleware.SwaggerUI(uiOpts, nil)
	router.Handle("/docs", swaggerUI)
	//
	// Documentation for sharing (See Redocyl, https://redocly.com/)
	redocOpts := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "redoc"}
	redocMiddleware := middleware.Redoc(redocOpts, nil)
	router.Handle("/redoc", redocMiddleware)
}

func server() {
	err := webServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start. Error is: %s", err.Error())
	}
}
