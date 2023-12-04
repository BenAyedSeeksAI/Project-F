package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
	database, _ := sql.Open("sqlite3", "./hrdb.db")
	return database
}
func EstablishDB() {
	db := OpenDB()
	defer db.Close()
	scriptFile := "./db/script.sql"
	sqlScript, err := ioutil.ReadFile(scriptFile)
	if err != nil {
		log.Fatal(err)
	}

	// Execute SQL script
	queries := string(sqlScript)
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SQL script executed successfully.")
}
