package utils

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

func GetTodayDate() time.Time {
	return time.Now()
}
func PrettyLogSuccess(message string) {
	fmt.Println("¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤")
	color.Cyan(message)
	fmt.Println("¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤")
}
func HasDependecies(dbse *sql.DB, tableName string, idColumnName string) (bool, error) {
	sqlStr := `SELECT COUNT(*)
	FROM pg_constraint c
	JOIN pg_attribute a ON a.attnum = ANY(c.conkey)
    WHERE confrelid = %s::regclass AND a.attname = %s`
	CompleteSqlStr := fmt.Sprintf(sqlStr, tableName, idColumnName)
	row := dbse.QueryRow(CompleteSqlStr)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return count > 0, nil

}
func IDExists(dbse *sql.DB, id uint, tableName string, idColumnName string) bool {
	sqlStr := `SELECT COUNT(*) FROM %s WHERE  %s= $1`
	CompleteSqlStr := fmt.Sprintf(sqlStr, tableName, idColumnName)
	row := dbse.QueryRow(CompleteSqlStr, id)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return count > 0
}
