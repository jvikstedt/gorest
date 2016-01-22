package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	fmt.Println("Starting gorest")

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", router)
}
