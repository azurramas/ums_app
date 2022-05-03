package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azurramas/ums_app/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	var users controllers.Users

	router := httprouter.New()

	// User Manipulation Endpoints
	router.GET("/users", users.List)
	router.POST("/user", users.Create)
	router.PUT("/user/:id", users.Edit)
	router.DELETE("/user/:id", users.Delete)

	fmt.Println("Starting server on port 3010...")
	log.Fatal(http.ListenAndServe(":3010", router))
}
