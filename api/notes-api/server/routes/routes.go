package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheussASM/pre-atividade-02/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		notes := main.Group("notes")
		{
			notes.GET("/", controllers.ShowAllNotes)
			notes.GET("/:id", controllers.ShowNote)
			notes.POST("/", controllers.CreateNote)
			notes.PUT("/", controllers.EditNote)
			notes.DELETE("/:id", controllers.DeleteNote)
		}
	}

	return router
}
