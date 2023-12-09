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
