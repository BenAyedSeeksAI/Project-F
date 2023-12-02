package db

import (
	"database/sql"

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
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	queries := []string{
		`DROP TABLE IF EXISTS Staff`,
		`DROP TABLE IF EXISTS Department`,
		`DROP TABLE IF EXISTS Candidate`,
		`CREATE TABLE Department (
			DepartmentID INTEGER PRIMARY KEY AUTOINCREMENT,
			DepartmentCode TEXT NOT NULL,
			DepartmentName TEXT NOT NULL,
			Location TEXT
			)`,
		`CREATE TABLE Staff (
			StaffID INTEGER PRIMARY KEY AUTOINCREMENT,
			FirstName TEXT NOT NULL,
			LastName TEXT NOT NULL,
			Email TEXT UNIQUE NOT NULL,
			Profession TEXT,
			Sex TEXT NOT NULL,
			DOB TEXT,
			Username TEXT,
			Password TEXT,
			PhoneNumber TEXT,
			HireDate TEXT,
			DepartmentID INT,
			CONSTRAINT FK_Staff_Department FOREIGN KEY (DepartmentID) REFERENCES Department(DepartmentID)
			)`,
		`CREATE TABLE Candidate (
            CandidateID INTEGER PRIMARY KEY AUTOINCREMENT,
			FirstName TEXT NOT NULL,
			LasttName TEXT NOT NULL,
			PersonalEmail TEXT NOT NULL,
			JobOffer TEXT NOT NULL,
			Degree TEXT NOT NULL,
			RecentExperience TEXT NOT NULL,
			Sex TEXT NOT NULL,
			DOB TEXT
		)`,
		`INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP1','Human Resources', 'Floor 1')`,
		`INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP2','IT (Cetec ERP)', 'Floor 1')`,
		`INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP3','Sales', 'Floor 1')`,
		`INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP4','Finance', 'Floor 2')`,
		`INSERT INTO Staff (FirstName, Lastname, Email, Profession, Sex, DOB, HireDate, DepartmentID) VALUES 
		('Mohamed Fares','Ben Ayed', 'feres.ba@qorevirtual.com', 'Software engineer', 'Male', '1997-03-06', '2023-01-20', 2)`,
	}
	for _, query := range queries {
		_, err := tx.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
