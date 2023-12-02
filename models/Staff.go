package models

import (
	"database/sql"
	"log"
	"time"
)

type Staff struct {
	StaffID        int       `json:"staff_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Profession     string    `json:"profession"`
	Sex            string    `json:"sex"`
	DOB            time.Time `json:"date_of_birth"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	PhoneNumber    string    `json:"phone_number"`
	HireDate       time.Time `json:"hire_date"`
	DepartmentID   int       `json:"department_id"`
	DepartmentCode string    `json:"department_code"`
	DepartmentName string    `json:"department_name"`
}

const timeLayout = "2006-01-02"

func DBInsertStaff(db *sql.DB, row *Staff) error {
	sqlStr := `INSERT INTO Staff (FirstName, LastName, Email, Profession, Sex, DOB, HireDate, DepartmentID) VALUES
	 (?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, row.FirstName, row.LastName, row.Email, row.Profession, row.Sex, row.DOB.Format(timeLayout), row.HireDate.Format(timeLayout), row.DepartmentID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func DBGetAllStaffDetails(db *sql.DB) ([]*Staff, error) {
	var staff []*Staff
	sqlStr := `SELECT s.StaffID, s.FirstName, s.LastName, s.Email, s.Profession, s.Sex, s.DOB, s.HireDate, s.DepartmentID, d.DepartmentCode, d.DepartmentName
	FROM Staff s
	JOIN Department d ON s.DepartmentID = d.DepartmentID`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err.Error())
		return staff, err
	}
	defer rows.Close()
	for rows.Next() {
		var Dob, HireDate string
		stf := new(Staff)
		err := rows.Scan(
			&stf.StaffID, &stf.FirstName, &stf.LastName,
			&stf.Email, &stf.Profession, &stf.Sex,
			&Dob, &HireDate, &stf.DepartmentID,
			&stf.DepartmentCode, &stf.DepartmentName,
		)
		if err != nil {
			log.Fatal(err)
			return []*Staff{}, err
		}
		if Dob != "" {
			stf.DOB, err = time.Parse(timeLayout, Dob)
			if err != nil {
				log.Fatal(err)
				return staff, err
			}
		}
		if HireDate != "" {
			stf.HireDate, err = time.Parse(timeLayout, HireDate)
			if err != nil {
				log.Fatal(err)
				return staff, err
			}
		}
		staff = append(staff, stf)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
		return staff, err
	}
	return staff, nil
}
