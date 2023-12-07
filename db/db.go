package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"rhmanager/utils"

	"log"

	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

const (
	configurationPath = "./configuration.json"
)

func OpenDB() *sql.DB {
	config := GetDBConfiguration(configurationPath)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
	database, _ := sql.Open("postgres", psqlInfo)
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
	SqlExecution(GetDBConfiguration(configurationPath).ScriptCreateFile, db)
	utils.PrettyLogSuccess("Database created Successfully ...")
	SqlExecution(GetDBConfiguration(configurationPath).ScriptSeedFile, db)
	utils.PrettyLogSuccess("Seed created Successfully ...")
	utils.PrettyLogSuccess("Database has been established ...")
}
