package main

import (
    "database/sql"
    
    _ "github.com/lib/pq"
    "fmt"
    //"time"
	"time"
)

var connection *sql.DB

func init() {
	connectionString := GetDatabaseConnectionString()

	var err error
	connection, err = sql.Open(
		"postgres",
		connectionString,
	)
	check(err)

	err = connection.Ping()
	check(err)
}

func TestConnect() {
    var pid int

    err := connection.QueryRow("SELECT pg_backend_pid()").Scan(&pid)
    check(err)
    fmt.Printf("pid: %v", pid)
}

func GetDatabaseConnectionString() string {
    connectionParams := GetFileContents(".connect")
    return connectionParams
}


func CreateTodo (todo *Todo) (error) {
	qryString := "INSERT INTO todo (name, due_date) VALUES ('%v', '%v'::DATE) RETURNING id, name, completed, due_date, entered_date, completed_date;"

	insertQry := fmt.Sprintf(qryString, todo.Name.String, todo.DueDate.Time.Format(time.RFC3339))

	err := connection.QueryRow(insertQry).Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.DueDate, &todo.EnteredDate, &todo.CompletedDate)

	return err
}

func GetAllTodos () ([]byte, error) {
	var result []byte
	selectAllQry := "SELECT json_agg(row_to_json(t.*)) FROM todo t;"

	err := connection.QueryRow(selectAllQry).Scan(&result)

	return result, err
}

func GetTodoByID (id int) (Todo, error) {
	var t Todo

	qryString := "SELECT * FROM todo WHERE id = %v"
	selectQry := fmt.Sprintf(qryString, id)

	err := connection.QueryRow(selectQry).Scan(&t.Id, &t.Name, &t.Completed, &t.DueDate, &t.EnteredDate, &t.CompletedDate)

	return t, err
}

func UpdateTodo(id int) error {
	// do stuff
	return nil
}

func DeleteTodo (id int) (Todo, error) {
	var t Todo

	qryString := "DELETE FROM todo WHERE id = %v RETURNING id, name, completed, due_date, entered_date, completed_date;"
	deleteQry := fmt.Sprintf(qryString, id)

	err := connection.QueryRow(deleteQry).Scan(&t.Id, &t.Name, &t.Completed, &t.DueDate, &t.EnteredDate, &t.CompletedDate)

	return t, err
}


