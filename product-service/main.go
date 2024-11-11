package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = make(map[string]Product)

func main() {
	http.HandleFunc("/products", createProduct)
	http.HandleFunc("/products/", getProduct)
	log.Println("Product Service listening on port 8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	products[product.ID] = product
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/products/"):]
	product, exists := products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}
