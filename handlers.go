package main

import (
	"fmt"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    
    Write200ResponseHeader(w, r)

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var _404 bool

    todoId, err := strconv.Atoi(vars["todoId"])
    
    if err != nil {        
        _404 = true
    }
    
    foundTodo, err := RepoFindTodo(todoId)
    
    if err != nil { 
        _404 = true
    }
    
    if !_404 {
        Write200ResponseHeader(w, r)
        SendAsJson(w, foundTodo)
    } else {
        Write404ResponseHeader(w, r)
    }
}
