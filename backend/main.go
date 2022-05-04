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
	router.GET("/perms", permissions.List)
	router.GET("/perms/:user_id", permissions.GetUserPerms)
	router.POST("/perms/:user_id/:perm_id", permissions.Add)
	router.DELETE("/perms/:user_id/:perm_id", permissions.Remove)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(services.CORS))
	n.UseHandler(router)

	url := ":" + app.Port
	n.Run(url)
}
