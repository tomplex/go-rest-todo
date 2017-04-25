package main

import "fmt"

var currentId int
var todos Todos

//set some seed data
func init() {
    RepoCreateTodo(Todo{Name: "finish this tutorial"})
    RepoCreateTodo(Todo{Name: "get a damn hairct"})
}

func RepoFindTodo(id int) (Todo, error) {
    for _, t := range todos {
        if t.Id == id {
            return t, nil
        }
    }
    return Todo{}, fmt.Errorf("Could not find Todo with ID of %d", id)
}

func RepoCreateTodo(t Todo) Todo {
    currentId += 1
    t.Id = currentId
    todos = append(todos, t)
    return t
}

func RepoDeleteTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with ID of %d to delete", id)
}
