package controller

import (
	"database/sql"
	"log"
	"net/http"
	"rhmanager/models"

	"github.com/gin-gonic/gin"
)

func InsertDepartment(c *gin.Context) {

	dbse, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
	}
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DBInsertDepartment(dbse.(*sql.DB), &department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert staff"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Department inserted successfully"})
}
func GetDepartmentDetails(c *gin.Context) {
	dbse, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
	}
	depmt, err := models.DBGetDepartmentDetails(dbse.(*sql.DB))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusOK, depmt)
}
