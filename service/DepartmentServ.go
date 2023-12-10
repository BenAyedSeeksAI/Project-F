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
func GetDepartmentByID(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	id := c.Param("id")
	IntId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department, err := models.DBGetDepartmentByID(dbse, uint(IntId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get department by id"})
		return
	}
	response.Success(c, http.StatusOK, department)
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
