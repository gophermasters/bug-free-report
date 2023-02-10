package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gophermasters/bug-free-report/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		Bugs := main.Group("bugs")
		{
			Bugs.GET("/", controllers.ShowAllBugsReports)
			Bugs.GET("/:id", controllers.ShowBugsReport)
			Bugs.POST("/", controllers.CreateBugsReport)
			Bugs.PUT("/", controllers.EditBugsReport)
			Bugs.DELETE("/:id", controllers.DeleteBugsReport)
		}
	}

	return router
}