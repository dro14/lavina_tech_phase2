package main

import (
	"Desktop/lavina_tech_phase2/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userId int = 32
var bookId int = 21

func main() {
	server := gin.Default()

	server.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.CreateNewUser(ctx, userId))
	})

	server.GET("/myself", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.GetUserInfo(ctx))
	})

	bookRoutes := server.Group("/books")
	{
		bookRoutes.POST("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, controller.CreateABook(ctx, bookId))
		})

		bookRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, controller.GetAllBooks())
		})

		bookRoutes.PATCH("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			ctx.JSON(http.StatusOK, controller.EditABook(ctx, bookId))
		})

		bookRoutes.DELETE("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			ctx.JSON(http.StatusOK, controller.DeleteABook(ctx, bookId))
		})
	}

	server.Run()
}