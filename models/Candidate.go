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
	sqlStr := `INSERT INTO Candidate (FirstName, LastName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB) VALUES
	 ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := db.Exec(sqlStr, row.FirstName, row.LastName, row.PersonalEmail, row.JobOffer, row.Degree, row.RecentExperience, row.Sex, row.DOB)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func DBGetAllCandidates(db *sql.DB) ([]*Candidate, error) {
	var candidate []*Candidate
	sqlStr := `SELECT CandidateID, FirstName, LastName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB 
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
			cand.DOB, err = time.Parse(time.RFC3339, dob)
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
	sqlStr := `SELECT FirstName, LastName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB 
	FROM Candidate 
	WHERE CandidateID = $1`
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
		candidateRow.DOB, err = time.Parse(time.RFC3339, dob)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return &candidateRow, nil
}
func DBDeleteCandidate(db *sql.DB, candID uint) error {
	sqlStr := `DELETE FROM Candidate WHERE CandidateID = $1`
	_, err := db.Exec(sqlStr, candID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func DBUpdateCandidate(db *sql.DB, candId uint, updatedCandidate Candidate) error {
	sqlStr := `UPDATE Candidate 
	SET FirstName= $1, LastName=$2 ,PersonalEmail=$3, JobOffer=$4, Degree=$5, RecentExperience=$6, Sex=$7, DOB=$8
	WHERE CandidateID=$9`
	_, err := db.Exec(sqlStr, updatedCandidate.FirstName, updatedCandidate.LastName,
		updatedCandidate.PersonalEmail, updatedCandidate.JobOffer,
		updatedCandidate.Degree, updatedCandidate.RecentExperience,
		updatedCandidate.Sex, updatedCandidate.DOB, candId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func DBHireCandidate(db *sql.DB, candId uint, StaffData CandidateToStaffData) error {
	candidateRow, err := DBGetCandidateByID(db, candId)
	if err != nil {
		return err
	}
	err = DBDeleteCandidate(db, candId)
	if err != nil {
		return err
	}
	stf := &Staff{
		FirstName:    candidateRow.FirstName,
		LastName:     candidateRow.LastName,
		Email:        StaffData.CompanyEmail,
		Profession:   candidateRow.JobOffer,
		Sex:          candidateRow.Sex,
		DOB:          candidateRow.DOB,
		HireDate:     utils.GetTodayDate(),
		DepartmentID: int(StaffData.DepartmentID),
	}
	err = DBInsertStaff(db, stf)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
