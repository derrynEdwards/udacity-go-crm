package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func getCustomers(w http.ResponseWriter, r *http.Request) {
	// Handles GET requests on "/customers" and returns
	// a JSON object list containing the Customer type structs.

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	// Handles GET requests on "/customers/{id}" and returns
	// a JSON object containing the Customer type struct.

	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	_, ok := Customers[id]

	if ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(getResponse("404"))
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	// Handles POST requests on "/customers" and returns
	// a JSON object containing the Customer type structs
	// if successfully created.

	w.Header().Set("Content-Type", "application/json")
	newEntry := map[string]Customer{}
	newCustomers := map[string]Customer{}
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &newEntry)

	for k, v := range newEntry {
		if _, ok := Customers[k]; ok {
			continue
		} else {
			newCustomers[k] = v
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(getResponse("400"))
	} else if len(newCustomers) > 0 {
		for k, v := range newCustomers {
			if v.Name != "" && v.Phone != "" {
				Customers[k] = v
			}
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Customers)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(getResponse("400"))
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Handles PUT requests on "/customers/{id}" and returns
	// a JSON object containing the updated Customer type struct.

	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	newEntry := Customer{}
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &newEntry)

	_, ok := Customers[id]

	if ok {
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(getResponse("400"))
		} else if newEntry.Name != "" && newEntry.Phone != "" {
			Customers[id] = newEntry
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Customers)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse("400"))
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(getResponse("404"))
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// Handles DELETE requests on "/customers/{id}" and returns
	// a JSON object containing the Customer type structs.

	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	_, ok := Customers[id]

	if ok {
		delete(Customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(getResponse("404"))
	}
}
