package routing

import (
	"database/sql"
	"net/http"
	"rhmanager/db"
	"rhmanager/service"

	"github.com/gin-gonic/gin"
)

func DatabaseMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
func RunRouting() {
	db := db.OpenDBForRoute()
	defer db.Close()
	r := gin.Default()
	r.Use(DatabaseMiddleware(db))
	r.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "Hello api"}) })
	departmentGroup := r.Group("/departments")
	{
		departmentGroup.GET("/get_all", service.GetDepartmentDetails)
		departmentGroup.POST("/create", service.InsertDepartment)
	}
	staffGroup := r.Group("/staffs")
	{
		staffGroup.GET("/get_all", service.GetStaffDetails)
		staffGroup.POST("/create", service.InsertStaff)
	}
	candidateGroup := r.Group("/candidates")
	{
		candidateGroup.GET("/get_all", service.GetCandidateDetails)
		candidateGroup.POST("/create", service.InsertCandidate)
		candidateGroup.POST("/delete", service.DeleteCandidate)
		candidateGroup.POST("/hire/:id/", service.HireCandidate)
	}
	r.Run(":8080")
}
