package controller

import (
	"database/sql"
	"net/http"
	"rhmanager/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertCandidate(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	var candidate models.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DBInsertCandidate(dbse, &candidate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert Candidate"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Candidate inserted successfully"})
}
func DeleteCandidate(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON body"})
		return
	}
	candId := uint(int(data["cand_id"].(float64)))

	err := models.DBDeleteCandidate(dbse, candId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete candidate"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Candidate: Successfull deletion!"})
}
func GetCandidateDetails(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	candidate, err := models.DBGetAllCandidates(dbse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all staff"})
		return
	}
	c.JSON(http.StatusOK, candidate)
}
func HireCandidate(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	var additionalStaffData models.CandidateToStaffData
	if err := c.ShouldBindJSON(&additionalStaffData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	IntId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DBHireCandidate(dbse, uint(IntId), additionalStaffData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all staff"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Candidate: Successfull Hiring to Staff!"})
}
