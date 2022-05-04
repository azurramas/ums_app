package main

import (
	"fmt"

	"github.com/azurramas/ums_app/controllers"
	"github.com/azurramas/ums_app/services"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {

	var users controllers.Users

	app := services.Init()

	fmt.Println("Connected to DB...")

	router := httprouter.New()

	// User Endpoints
	router.GET("/users", users.List)
	router.GET("/user/:id", users.Get)
	router.POST("/user", users.Create)
	router.PUT("/user", users.Edit)
	router.DELETE("/user/:id", users.Delete)

	n := negroni.Classic()
	n.UseHandler(router)

	url := ":" + app.Port
	n.Run(url)
}
