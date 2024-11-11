package main

import (
    "io"
    "log"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func proxyRequest(w http.ResponseWriter, r *http.Request) {
    var targetURL string

    // Determine the target URL based on the path
    if strings.HasPrefix(r.URL.Path, "/user") || strings.HasPrefix(r.URL.Path, "/register") {
        targetURL = "http://localhost:8080" + r.URL.Path // User Service
	} else if strings.HasPrefix(r.URL.Path, "/order") || strings.HasPrefix(r.URL.Path, "/orders") {
			targetURL = "http://localhost:8081" + r.URL.Path // Order Service
    } else {
        http.Error(w, "Service not found", http.StatusNotFound)
        return
    }

    // Create a new request with the target URL and the same method and body as the original request
    req, err := http.NewRequest(r.Method, targetURL, r.Body)
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }

    // Copy the original request headers to the new request
    req.Header = r.Header.Clone()

    // Forward the request to the target service
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
        return
    }
    defer resp.Body.Close()

    // Copy the response headers and status code from the target service
    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)

    // Copy the response body from the target service to the client
    io.Copy(w, resp.Body)
}

func main() {
    r := mux.NewRouter()
    r.PathPrefix("/").HandlerFunc(proxyRequest)

    handler := corsMiddleware(r)
    log.Println("API Gateway running on port 8082...")
    log.Fatal(http.ListenAndServe(":8082", handler))
}
