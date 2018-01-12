package main

import (
	"fmt"
    "encoding/json"
    "net/http"
    "strconv"
    "io"
    "io/ioutil"

    "github.com/gorilla/mux"
)

var HeaderKey = "Content-Type"
var JSONContent = "application/json; charset=UTF-8"


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    
    Write200ResponseHeader(w, r)

	todos, err := GetAllTodos()
	check(err)

	w.Write(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var _404 bool

    todoId, err := strconv.Atoi(vars["todoId"])

    if err != nil {        
        _404 = true
    }
    
    foundTodo, err := GetTodoByID(todoId)
    
    if err != nil { 
        _404 = true
    }
    
    if !_404 {
        Write200ResponseHeader(w, r)
		json.NewEncoder(w).Encode(foundTodo)
    } else {
        Write404ResponseHeader(w, r)
    }
}

func TodoCreate (w http.ResponseWriter, r *http.Request) {
    var todo Todo

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    check(err)

    check(r.Body.Close())

	w.Header().Set(HeaderKey, JSONContent)

    if err := json.Unmarshal(body, &todo); err != nil {
        // We weren't able to successfully unmarshal the JSON that was POST-ed
        w.WriteHeader(422)
        err := json.NewEncoder(w).Encode(err)
        checkWithMsg(err, "Couldn't send error message for failed create request")
    }
    // we were, woo
	err = CreateTodo(&todo)
	checkWithMsg(err, "couldn't create todo")

    w.WriteHeader(http.StatusCreated)
    // respond with the JSON of the new todo
    err = json.NewEncoder(w).Encode(todo)
    checkWithMsg(err, "couldn't write header")
}

func TodoDelete (w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	check(err)

	w.Header().Set(HeaderKey, JSONContent)

	if err := json.Unmarshal(body, &todo); err != nil {
		w.WriteHeader(422)
		err := json.NewEncoder(w).Encode(err)
		checkWithMsg(err, "Couldn't send error message for failed delete request")
	}

	deletedTodo, err := DeleteTodo(int(todo.Id.Int64))
	check(err)

	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(deletedTodo)
	checkWithMsg(err, "Couldn't write header for delete request")
	//vars := mux.Vars(r)
	//
	//todoId, err := strconv.Atoi(vars["todoId"])
	//
	//if err != nil {
	//	Write404ResponseHeader(w, r)
	//	return
	//}
	// deletedTodo, err := DeleteTodo(todoId)
	//
	//if err != nil {
	//	Write404ResponseHeader(w, r)
	//	return
	//} else {
	//	Write200ResponseHeader(w, r)
	//	json.NewEncoder(w).Encode(deletedTodo)
	//}
}