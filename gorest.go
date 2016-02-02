package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/jvikstedt/gorest/controllers"
)

func main() {
	fmt.Println("Starting gorest")
	router := httprouter.New()

	uc := controllers.NewUserController(getSession())

	router.GET("/users/:id", uc.GetUser)
	router.GET("/users", uc.GetUsers)
	router.POST("/users", uc.CreateUser)
	router.DELETE("/users/:id", uc.DeleteUser)
	router.PATCH("/users/:id", uc.UpdateUser)

	http.ListenAndServe(":3000", router)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
