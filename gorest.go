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
	router.POST("/users", uc.CreateUser)

	http.ListenAndServe(":3000", router)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
