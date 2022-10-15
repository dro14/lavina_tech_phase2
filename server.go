package main

import (
	"Desktop/lavina_tech_phase2/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.CreateNewUser(ctx))
	})

	server.GET("/myself", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.GetUserInfo(ctx))
	})

	bookRoutes := server.Group("/books")
	{
		bookRoutes.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, controller.CreateABook(ctx))
		})

		bookRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, controller.GetAllBooks(ctx))
		})

		bookRoutes.PATCH("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, controller.EditABook(ctx))
		})

		bookRoutes.DELETE("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, controller.DeleteABook(ctx))
		})
	}

	server.Run()
}