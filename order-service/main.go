package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"

    "github.com/gorilla/mux"
)

// Order struct representing an order
type Order struct {
    Username string `json:"username"`
    Product  string `json:"product"`
    Quantity int    `json:"quantity"`
}

// Store for orders (in-memory for simplicity)
var orders []Order
var mu sync.Mutex // to handle concurrent access to orders slice

// PlaceOrder endpoint to create a new order
func placeOrder(w http.ResponseWriter, r *http.Request) {
    var order Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        http.Error(w, "Invalid order data", http.StatusBadRequest)
        return
    }

    mu.Lock()
    orders = append(orders, order)
    mu.Unlock()

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Order placed successfully"))
}

// ViewOrders endpoint to get all orders for a specific user
func viewOrders(w http.ResponseWriter, r *http.Request) {
    username := mux.Vars(r)["username"]
    userOrders := []Order{}

    mu.Lock()
    for _, order := range orders {
        if order.Username == username {
            userOrders = append(userOrders, order)
        }
    }
    mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userOrders)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/order", placeOrder).Methods("POST")
    r.HandleFunc("/orders/{username}", viewOrders).Methods("GET")

    log.Println("Order Service running on port 8081...")
    log.Fatal(http.ListenAndServe(":8081", r))
}
