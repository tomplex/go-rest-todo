package main

import (
    "database/sql"
    
    _ "github.com/lib/pq"
    "fmt"
)

func TestConnect() {
    connectionString := GetDatabaseConnectionString()

    db, err := sql.Open(
        "postgres",
        connectionString,
        )
    check(err)
    
    var pid int
    err = db.QueryRow("SELECT pg_backend_pid()").Scan(&pid)
    check(err)
    fmt.Printf("pid: %v", pid)
}

func GetDatabaseConnectionString() string {
    connectionParams := GetFileContents(".connect")
    return connectionParams
}


//func CreateTodo (t Todo) {}
//
//func GetTodo (id int) {}
//
//func UpdateTodo(id int) {}
//
//func DeleteTodo () {}


