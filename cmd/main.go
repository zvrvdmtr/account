package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zvrvdmtr/account/pkg/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", api.Ping).Methods("GET")
	router.HandleFunc("/user/{id}/balance", api.GetUserAccountBalance()).Methods("GET")
	router.HandleFunc("/user/{id}/balance/change", api.ChangeUserAccountBalance()).Methods("POST")
	router.HandleFunc("/transaction", api.TransactionUserToUser()).Methods("POST")
	fmt.Println("Server run on port 8000")
	http.ListenAndServe(":8000", router)
}
