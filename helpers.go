package main

import (
    "net/http"
    "encoding/json"
)

func Write404ResponseHeader(w http.ResponseWriter, r *http.Request) { 
    http.NotFound(w, r)
}

func Write200ResponseHeader(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func SendAsJson(w http.ResponseWriter, t Todo) {
    json.NewEncoder(w).Encode(t)
}
