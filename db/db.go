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

func OpenDBForRoute() *sql.DB {
	config := GetDBConfiguration(configurationPath)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
	database, _ := sql.Open("postgres", psqlInfo)
	return database
}
func OpenDB() (*sql.DB, *Config) {
	config := GetDBConfiguration(configurationPath)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
	database, _ := sql.Open("postgres", psqlInfo)
	return database, config
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
	db, DBConfig := OpenDB()
	defer db.Close()
	SqlExecution(DBConfig.ScriptCreateFile, db)
	utils.PrettyLogSuccess("Database created Successfully ...")
	SqlExecution(DBConfig.ScriptSeedFile, db)
	utils.PrettyLogSuccess("Seed created Successfully ...")
	utils.PrettyLogSuccess("Database has been established ...")
}
