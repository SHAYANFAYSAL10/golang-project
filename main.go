package main

import (
	"VIVASOFT1/src/apis/user_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/findall", user_api.FindAll).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println((err))
	}
}
