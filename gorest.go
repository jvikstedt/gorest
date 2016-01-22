package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vikstedt/gorest/controllers"
)

func main() {
	fmt.Println("Starting gorest")
	router := httprouter.New()

	uc := controllers.NewUserController()

	router.GET("/users/:id", uc.GetUser)

	http.ListenAndServe(":3000", router)
}
