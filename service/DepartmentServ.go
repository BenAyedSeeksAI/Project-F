package service

import (
	"database/sql"
	"log"
	"net/http"
	"rhmanager/models"
	"rhmanager/response"

	"github.com/gin-gonic/gin"
)

func InsertDepartment(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.DBInsertDepartment(dbse, &department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert staff"})
		return
	}
	response.SuccessOperation(c, http.StatusCreated, "InsertDepartment", response.INSERTION)
}
func GetDepartmentDetails(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	depmt, err := models.DBGetDepartmentDetails(dbse)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	response.Success(c, http.StatusOK, depmt)
}
