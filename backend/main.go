package main

import (
	"fmt"

	"github.com/azurramas/ums_app/controllers"
	"github.com/azurramas/ums_app/services"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {

	var (
		users       controllers.Users
		permissions controllers.Permissions
	)

	app := services.Init()

	fmt.Println("Connected to DB...")

	router := httprouter.New()

	// User Endpoints
	router.GET("/users", users.List)
	router.GET("/user/:id", users.Get)
	router.POST("/user", users.Create)
	router.PUT("/user", users.Edit)
	router.DELETE("/user/:id", users.Delete)

	// Permission Endpoints
	router.POST("/perms", permissions.Add)
	router.DELETE("/perms/:id", permissions.Remove)

	n := negroni.Classic()
	n.UseHandler(router)

	url := ":" + app.Port
	n.Run(url)
}
