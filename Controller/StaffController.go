package controller

import (
	"database/sql"
	"log"
	"net/http"
	"rhmanager/models"

	"github.com/gin-gonic/gin"
)

func InsertStaff(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	var staff models.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.DBInsertStaff(dbse, &staff)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert staff"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Staff inserted successfully"})
}
func GetStaffDetails(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	staff, err := models.DBGetAllStaffDetails(dbse)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusOK, staff)
}
