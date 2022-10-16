package main

import (
	"net/http"

	"github.com/dro14/lavina_tech_phase2/controller"

	"github.com/gin-gonic/gin"
)

const URL string = "http://localhost:5000"
var tempUrl string = URL
var userId int = 1
var bookId int = 1

const messageFor401 string = "Unauthorized, permission denied. Enter the following credentials as headers: Key: {Key}, Sign: {Sign}"


func main() {
	server := gin.Default()

	server.POST("/signup", func(ctx *gin.Context) {
		tempUrl += "/signup"
		ctx.JSON(http.StatusOK, controller.CreateNewUser(ctx, userId))
		userId++
		tempUrl = URL
	})

	server.GET("/myself", func(ctx *gin.Context) {
		tempUrl += "/myself"
		if controller.Authorize(ctx, "GET", tempUrl) {
			ctx.JSON(http.StatusOK, controller.GetUserInfo(ctx))
		} else {
			// request, _ := ioutil.ReadAll(ctx.Request.Body)
			ctx.JSON(http.StatusUnauthorized, messageFor401)
		}
		tempUrl = URL
	})

	bookRoutes := server.Group("/books")
	{
		bookRoutes.POST("", func(ctx *gin.Context) {
			tempUrl += "/books"
			if controller.Authorize(ctx, "POST", tempUrl) {
				ctx.JSON(http.StatusOK, controller.CreateABook(ctx, bookId))
				bookId++
			} else {
				// request, _ := ioutil.ReadAll(ctx.Request.Body)
				ctx.JSON(http.StatusUnauthorized, messageFor401)
			}
			tempUrl = URL
		})

		bookRoutes.GET("", func(ctx *gin.Context) {
			tempUrl += "/books"
			if controller.Authorize(ctx, "GET", tempUrl) {
				ctx.JSON(http.StatusOK, controller.GetAllBooks())
			} else {
				// request, _ := ioutil.ReadAll(ctx.Request.Body)
				ctx.JSON(http.StatusUnauthorized, messageFor401)
			}
			tempUrl = URL
		})

		bookRoutes.PATCH("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			tempUrl += "/books/" + bookId
			if controller.Authorize(ctx, "PATCH", tempUrl) {
				ctx.JSON(http.StatusOK, controller.EditABook(ctx, bookId))
			} else {
				// request, _ := ioutil.ReadAll(ctx.Request.Body)
				ctx.JSON(http.StatusUnauthorized, messageFor401)
			}
			tempUrl = URL
		})

		bookRoutes.GET("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			tempUrl += "/books/" + bookId
			if controller.Authorize(ctx, "GET", tempUrl) {
				ctx.JSON(http.StatusOK, controller.DeleteABook(ctx, bookId))
			} else {
				// request, _ := ioutil.ReadAll(ctx.Request.Body)
				ctx.JSON(http.StatusUnauthorized, messageFor401)
			}
			tempUrl = URL
		})
	}

	server.Run()
}