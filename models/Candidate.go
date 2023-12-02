package models

import (
	"database/sql"
	"log"
	"rhmanager/utils"
	"time"
)

type Candidate struct {
	CandidateID      int       `json:"candidate_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	PersonalEmail    string    `json:"personal_email"`
	JobOffer         string    `json:"job_offer"`
	Degree           string    `json:"degree"`
	RecentExperience string    `json:"recent_experience"`
	Sex              string    `json:"sex"`
	DOB              time.Time `json:"date_of_birth"`
}

func DBInsertCandidate(db *sql.DB, row *Candidate) error {
	sqlStr := `INSERT INTO Candidate (FirstName, LasttName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB) VALUES
	 (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStr, row.FirstName, row.LastName, row.PersonalEmail, row.JobOffer, row.Degree, row.RecentExperience, row.Sex, row.DOB.Format(timeLayout))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func DBGetAllCandidates(db *sql.DB) ([]*Candidate, error) {
	var candidate []*Candidate
	sqlStr := `SELECT CandidateID, FirstName, LasttName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB 
	FROM Candidate`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err.Error())
		return []*Candidate{}, err
	}
	defer rows.Close()
	for rows.Next() {
		cand := new(Candidate)
		var dob string
		err := rows.Scan(
			&cand.CandidateID, &cand.FirstName, &cand.LastName,
			&cand.PersonalEmail, &cand.JobOffer, &cand.Degree,
			&cand.RecentExperience, &cand.Sex, &dob,
		)
		if err != nil {
			log.Fatal(err)
			return candidate, err
		}
		if dob != "" {
			cand.DOB, err = time.Parse(timeLayout, dob)
			if err != nil {
				log.Fatal(err)
				return candidate, err
			}
		}
		candidate = append(candidate, cand)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
		return candidate, err
	}
	return candidate, nil
}

func DBGetCandidateByID(db *sql.DB, candId uint) (*Candidate, error) {
	sqlStr := `SELECT FirstName, LasttName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB 
	FROM Candidate 
	WHERE CandidateID = ?`
	row := db.QueryRow(sqlStr, candId)

	var candidateRow Candidate
	var dob string
	err := row.Scan(
		&candidateRow.FirstName, &candidateRow.LastName, &candidateRow.PersonalEmail,
		&candidateRow.JobOffer, &candidateRow.Degree, &candidateRow.RecentExperience,
		&candidateRow.Sex, &dob,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if dob != "" {
		candidateRow.DOB, err = time.Parse(timeLayout, dob)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return &candidateRow, nil
}
func DBDeleteCandidate(db *sql.DB, candId uint) error {
	sqlStr := `DELETE FROM Candidate WHERE CandidateID=? `
	_, err := db.Exec(sqlStr, candId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DBHireCandidate(db *sql.DB, candId uint, departmentId uint) error {
	candidateRow, err := DBGetCandidateByID(db, candId)
	if candidateRow == nil || err != nil {
		return err
	}
	err = DBDeleteCandidate(db, candId)
	if err != nil {
		return err
	}
	stf := &Staff{
		FirstName:    candidateRow.FirstName,
		LastName:     candidateRow.LastName,
		Email:        "default@default.xom",
		Profession:   candidateRow.JobOffer,
		Sex:          candidateRow.Sex,
		DOB:          candidateRow.DOB,
		HireDate:     utils.GetTodayDate(),
		DepartmentID: int(departmentId),
	}
	err = DBInsertStaff(db, stf)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
