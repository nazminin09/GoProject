package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Item represents a simple item in our API
type Item struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Initialize an in-memory database of items
var items = map[string]Item{}

// Handle GET requests to retrieve all items
func getItems(w http.ResponseWriter, r *http.Request) {
	itemsList := make([]Item, 0, len(items))
	for _, v := range items {
		itemsList = append(itemsList, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemsList)
}

// Handle GET requests to retrieve a specific item
func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID := params["id"]

	if item, ok := items[itemID]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Item with ID %s not found", itemID)
	}
}

// Handle POST requests to create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)

	itemID := fmt.Sprintf("%d", len(items)+1)
	newItem.ID = itemID

	items[itemID] = newItem

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

func main() {
	// Create a new router instance
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")

	// Start the server
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
