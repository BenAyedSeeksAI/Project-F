package models

import (
	"database/sql"
	"log"
)

type Department struct {
	DepartmentID       int    `json:"dp_id"`
	DepartmentCode     string `json:"dp_code"`
	DepartmentName     string `json:"dp_name"`
	DepartmentLocation string `json:"dp_location"`
}

func DBInsertDepartment(db *sql.DB, row *Department) error {
	sqlStr := `INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES
	 ($1,$2,$3)`
	_, err := db.Exec(sqlStr, row.DepartmentCode, row.DepartmentName, row.DepartmentLocation)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func DBGetDepartmentDetails(db *sql.DB) ([]*Department, error) {
	var depts []*Department
	sqlStr := `SELECT DepartmentID, DepartmentCode, DepartmentName, Location FROM Department`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err.Error())
		return depts, err
	}
	defer rows.Close()
	for rows.Next() {
		depmt := new(Department)
		if err := rows.Scan(
			&depmt.DepartmentID, &depmt.DepartmentCode, &depmt.DepartmentName, &depmt.DepartmentLocation,
		); err != nil {
			log.Fatal(err.Error())
			return depts, err
		}
		depts = append(depts, depmt)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
		return depts, err
	}
	return depts, nil
}
