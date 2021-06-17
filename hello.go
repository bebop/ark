package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello, World!")
	database, _ := sql.Open("sqlite3", "./allbase.db")
	rows, _ := database.Query("SELECT name, formula FROM compound")
	var name string
	var formula string
	for rows.Next() {
		rows.Scan(&name, &formula)
		fmt.Println(name, ": ", formula)
	}

}
