package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID     string `json:"ID"`
	Title  string `json:"Title"`
	Year   int64  `json:"Year"`
	Author string `json:"Author"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Function Called: homePage()")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: allArticles()")

	json.NewEncoder(w).Encode(Articles)
}

func newArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article) // Obtain item from request JSON

	Articles = append(Articles, article) // Add item to inventory

	json.NewEncoder(w).Encode(article) // Show item in response JSON for verification
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtId(params["id"])

	json.NewEncoder(w).Encode(Articles)
}

func _deleteItemAtId(id string) {
	for index, Article := range Articles {
		if Article.ID == id {
			// Delete item from Slice
			Articles = append(Articles[:index], Articles[index+1:]...)
			break
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article) // Obtain item from request JSON

	params := mux.Vars(r)

	_deleteItemAtId(params["id"])        // Delete item
	Articles = append(Articles, article) // Create it again with data from request

	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	// := is the short variable declaration operator
	// Automatically determines type for variable
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/allArticles", allArticles).Methods("GET")
	router.HandleFunc("/allArticles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/allArticles/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/allArticles", newArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	// Data store
	Articles = append(Articles, Article{
		ID:     "0",
		Title:  "Golang Basics",
		Year:   2020,
		Author: "Sumeet",
	})

	handleRequests()
}
