package main

import (
	"github.com/dro14/lavina_tech_phase2/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

const URL string = "https://lavina-tech-phase2.herokuapp.com"
var tempUrl string = URL
var userId int = 1
var bookId int = 1

func main() {
	server := gin.Default()

	server.POST("/signup", func(ctx *gin.Context) {
		tempUrl += "/signup"
		if controller.Authorize("POST", tempUrl) {
			ctx.JSON(http.StatusOK, controller.CreateNewUser(ctx, userId))
			userId++
		} else {
			ctx.JSON(http.StatusUnauthorized,
				"Unauthorized, permission denied.\n" +
				"Enter the following credentials as headers:\n" +
				"Key: {Key}\n" +
				"Sign: {Sign}\n")
		}
		tempUrl = URL
	})

	server.GET("/myself", func(ctx *gin.Context) {
		tempUrl += "/myself"
		if controller.Authorize("GET", tempUrl) {
			ctx.JSON(http.StatusOK, controller.GetUserInfo(ctx))
		} else {
			ctx.JSON(http.StatusUnauthorized,
				"Unauthorized, permission denied.\n" +
				"Enter the following credentials as headers:\n" +
				"Key: {Key}\n" +
				"Sign: {Sign}\n")
		}
		tempUrl = URL
	})

	bookRoutes := server.Group("/books")
	{
		bookRoutes.POST("", func(ctx *gin.Context) {
			tempUrl += "/books"
			if controller.Authorize("POST", tempUrl) {
				ctx.JSON(http.StatusOK, controller.CreateABook(ctx, bookId))
				bookId++
			} else {
				ctx.JSON(http.StatusUnauthorized,
					"Unauthorized, permission denied.\n" +
					"Enter the following credentials as headers:\n" +
					"Key: {Key}\n" +
					"Sign: {Sign}\n")
			}
			tempUrl = URL
		})

		bookRoutes.GET("", func(ctx *gin.Context) {
			tempUrl += "/books"
			if controller.Authorize("GET", tempUrl) {
				ctx.JSON(http.StatusOK, controller.GetAllBooks())
			} else {
				ctx.JSON(http.StatusUnauthorized,
					"Unauthorized, permission denied.\n" +
					"Enter the following credentials as headers:\n" +
					"Key: {Key}\n" +
					"Sign: {Sign}\n")
			}
			tempUrl = URL
		})

		bookRoutes.PATCH("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			tempUrl += "/books/" + bookId
			if controller.Authorize("PATCH", tempUrl) {
				ctx.JSON(http.StatusOK, controller.EditABook(ctx, bookId))
			} else {
				ctx.JSON(http.StatusUnauthorized,
					"Unauthorized, permission denied.\n" +
					"Enter the following credentials as headers:\n" +
					"Key: {Key}\n" +
					"Sign: {Sign}\n")
			}
			tempUrl = URL
		})

		bookRoutes.DELETE("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			tempUrl += "/books/" + bookId
			if controller.Authorize("PATCH", tempUrl) {
				ctx.JSON(http.StatusOK, controller.DeleteABook(ctx, bookId))
			} else {
				ctx.JSON(http.StatusUnauthorized,
					"Unauthorized, permission denied.\n" +
					"Enter the following credentials as headers:\n" +
					"Key: {Key}\n" +
					"Sign: {Sign}\n")
			}
			tempUrl = URL
		})
	}

	server.Run()
}