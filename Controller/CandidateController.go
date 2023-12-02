package controller

import (
	"database/sql"
	"net/http"
	"rhmanager/models"

	"github.com/gin-gonic/gin"
)

func InsertCandidate(c *gin.Context) {

	dbse, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
	}
	var candidate models.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DBInsertCandidate(dbse.(*sql.DB), &candidate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert Candidate"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Candidate inserted successfully"})
}

func GetCandidateDetails(c *gin.Context) {
	dbse, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
	}
	candidate, err := models.DBGetAllCandidates(dbse.(*sql.DB))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all staff"})
		return
	}
	c.JSON(http.StatusOK, candidate)
}
