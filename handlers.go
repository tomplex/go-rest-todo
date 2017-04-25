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
    //todos := Todos{
        //Todo{Name: "Finish this tutorial"},
        //Todo{Name: "Get a haircut"},
    //}
    
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    todoId, err := strconv.Atoi(vars["todoId"])
    
    if err != nil {        
        w.WriteHeader(404)
        http.NotFound(w, r)
    }
    
    if err := json.NewEncoder(w).Encode(RepoFindTodo(todoId)); err != nil {
        panic(err)
    }
}
