package service

import (
	"database/sql"
	"log"
	"net/http"
	"rhmanager/models"
	"rhmanager/response"
	"rhmanager/utils"
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
func DeleteDepartment(c *gin.Context) {
	dbse := c.MustGet("db").(*sql.DB)
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON body"})
		return
	}
	ok, err := utils.HasDependecies(dbse, "Department", "DepartmentID")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete candidate"})
		return
	}
	if ok {
		deptID := uint(int(data["dept_id"].(float64)))
		err = models.DBDeleteCandidate(dbse, deptID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete candidate"})
			return
		}
		response.SuccessOperation(c, http.StatusCreated, "InsertDepartment", response.DELETION)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "No deletion accured due to db constraints!"})
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
