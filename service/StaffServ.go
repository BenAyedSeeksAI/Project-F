package service

import (
	"database/sql"
	"log"
	"net/http"
	"rhmanager/models"
	"rhmanager/response"
	"strconv"

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
	response.SuccessOperation(c, http.StatusCreated, "InsertStaff", response.INSERTION)
}
func GetStaffDetails(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	staff, err := models.DBGetAllStaffDetails(dbse)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	response.Success(c, http.StatusOK, staff)
}
func GetStaffByID(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	id := c.Param("id")
	IntId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	staff, err := models.DBGetStaffByID(dbse, uint(IntId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, http.StatusOK, staff)
}
