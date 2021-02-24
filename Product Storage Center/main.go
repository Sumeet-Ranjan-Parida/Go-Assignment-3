package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	PID      string `json:"PID"`
	Name     string `json:"Name"`
	Desc     string `json:"Desc"`
	Quantity int64  `json:"Quantity"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Function Called: homePage()")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getInventory()")

	json.NewEncoder(w).Encode(inventory)
}

func newProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtPid(params["pid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtPid(pid string) {
	for index, item := range inventory {
		if item.PID == pid {

			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	_deleteItemAtPid(params["pid"])
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(inventory)
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getProducts).Methods("GET")
	router.HandleFunc("/inventory/{uid}", updateProduct).Methods("PUT")
	router.HandleFunc("/inventory/{uid}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/inventory", newProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {

	inventory = append(inventory, Item{
		PID:      "0",
		Name:     "Apple MacBook Pro",
		Desc:     "MacBook Pro 13-inch",
		Quantity: 1,
	})

	handleRequests()
}
