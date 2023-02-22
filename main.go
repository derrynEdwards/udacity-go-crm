package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Initial Database
var Customers = map[string]Customer{
	"1": {"John", "CEO", "john@udatech.com", "+11234567890", false},
	"2": {"Jane", "COO", "jane@udatech.com", "+11234567908", true},
	"3": {"Carl", "CEO", "carl@mycompany.com", "+11234567809", true},
	"4": {"Kane", "CTO", "kane@kanecorp.com", "+11234567788", false},
}

func main() {
	router := mux.NewRouter()

	//Routes and Handlers
	router.Handle("/", http.FileServer(http.Dir("./static"))).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting server on 0.0.0.0:3000...")

	http.ListenAndServe(":3000", router)
}
