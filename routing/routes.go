package routing

import (
	"database/sql"
	"net/http"
	controller "rhmanager/Controller"
	"rhmanager/db"

	"github.com/gin-gonic/gin"
)

func DatabaseMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
func RunRouting() {
	db := db.OpenDB()
	r := gin.Default()
	r.Use(DatabaseMiddleware(db))
	r.Use(gin.Logger())
	r.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "Hello api"}) })
	departmentGroup := r.Group("/departments")
	{
		departmentGroup.GET("/get_all", controller.GetDepartmentDetails)
		departmentGroup.POST("/create", controller.InsertDepartment)
	}
	staffGroup := r.Group("/staffs")
	{
		staffGroup.GET("/get_all", controller.GetStaffDetails)
		staffGroup.POST("/create", controller.InsertStaff)
	}
	candidateGroup := r.Group("/candidates")
	{
		candidateGroup.POST("/create", controller.InsertCandidate)
		candidateGroup.GET("/get_all", controller.GetCandidateDetails)
	}
	r.Run(":8080")
}
