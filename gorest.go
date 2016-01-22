package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	fmt.Println("Starting gorest")

	router := httprouter.New()
	router.GET("/", Index)

	http.ListenAndServe(":3000", router)
}
