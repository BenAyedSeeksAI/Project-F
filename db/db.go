package db

import (
	"database/sql"
	"io/ioutil"
	"rhmanager/utils"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	scriptCreateFile = "./db/script_create.sql"
	scriptSeedFile   = "./db/script_seed.sql"
)

func OpenDB() *sql.DB {
	database, _ := sql.Open("sqlite3", "./hrdb.db")
	return database
}
func SqlExecution(scriptPath string, db *sql.DB) {
	sqlScript, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		log.Fatal(err)
	}
	queries := string(sqlScript)
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
}
func EstablishDB() {
	db := OpenDB()
	defer db.Close()
	SqlExecution(scriptCreateFile, db)
	utils.PrettyLogSuccess("Database created Successfully ...")
	SqlExecution(scriptSeedFile, db)
	utils.PrettyLogSuccess("Seed created Successfully ...")
	utils.PrettyLogSuccess("Database has been established ...")
}
