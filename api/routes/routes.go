package routes

import (
	"sqlc/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthorRoutes(router *gin.Engine, authorController *controllers.AuthorController) {
	router.GET("/authors", authorController.ListAuthors)
	router.POST("/authors", authorController.CreateAuthor)
	// router.GET("/authors/:id", controllers.GetAuthor)
	// router.PUT("/authors/:id", controllers.UpdateAuthor)
	// router.DELETE("/authors/:id", controllers.DeleteAuthor)
}
