package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/swhite24/go-rest-tutorial/controllers"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a user resource
	r.GET("/user/:id", controllers.GetUser)

	r.POST("/user", controllers.CreateUser)

	r.DELETE("/user/:id", controllers.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
