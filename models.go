package main

import (
	"database/sql"
	//"github.com/lib/pq"
	"strconv"
)


type Todos []Todo


type Todo struct {
	Id          JsonNullInt64     `json:"id"`
	Name        JsonNullString	`json:"name"`
	Completed   JsonNullBool	`json:"completed"`
	DueDate         JsonNullTime	`json:"due_date"`
	EnteredDate		JsonNullTime	`json:"entered_date"`
	CompletedDate	JsonNullTime	`json:"completed_date"`
}

//ToNullString invalidates a sql.NullString if empty, validates if not empty
func ToNullString(s string) sql.NullString {
	return sql.NullString{String : s, Valid : s != ""}
}


//ToNullInt64 validates a sql.NullInt64 if incoming string evaluates to an integer, invalidates if it does not
func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64 : int64(i), Valid : err == nil}
}

