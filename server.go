package main

import (
	"net/http"

	"Desktop/lavina_tech_phase2/controller"
	"Desktop/lavina_tech_phase2/service"

	"github.com/gin-gonic/gin"
)

const hostURL string = "http://localhost:5000"

var (
	userId      int                   = 1
	bookId      int                   = 1
	tempUrl     string                = hostURL
	userService service.UserService   = service.NewUserList()
	bookService service.BookService   = service.NewBookList()
	_controller controller.Controller = controller.NewController(userService, bookService)
)

func main() {
	var server *gin.Engine = gin.Default()

	// POST Create new user
	server.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, _controller.CreateNewUser(ctx, &userId))
	})

	// GET Get user info
	server.GET("/myself", func(ctx *gin.Context) {
		tempUrl += "/myself"
		isAuthorized, _ := _controller.AuthorizeUser(ctx, "GET", tempUrl)
		if isAuthorized {
			ctx.JSON(http.StatusOK, _controller.GetUserInfo(ctx))
		}
		tempUrl = hostURL
	})

	var bookRoutes *gin.RouterGroup = server.Group("/books")
	{
		// POST Create a book
		bookRoutes.POST("", func(ctx *gin.Context) {
			tempUrl += "/books"
			isAuthorized, body := _controller.AuthorizeUser(ctx, "POST", tempUrl)
			if isAuthorized {
				ctx.JSON(http.StatusOK, _controller.CreateABook(ctx, &bookId, body))
			}
			tempUrl = hostURL
		})

		// GET Get all books
		bookRoutes.GET("", func(ctx *gin.Context) {
			tempUrl += "/books"
			isAuthorized, _ := _controller.AuthorizeUser(ctx, "GET", tempUrl)
			if isAuthorized {
				ctx.JSON(http.StatusOK, _controller.GetAllBooks(ctx))
			}
			tempUrl = hostURL
		})

		// PATCH Edit a book
		bookRoutes.PATCH("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			tempUrl += "/books/" + bookId
			isAuthorized, body := _controller.AuthorizeUser(ctx, "PATCH", tempUrl)
			if isAuthorized {
				ctx.JSON(http.StatusOK, _controller.EditABook(ctx, bookId, body))
			}
			tempUrl = hostURL
		})

		// GET Delete a book
		bookRoutes.DELETE("/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			tempUrl += "/books/" + bookId
			isAuthorized, _ := _controller.AuthorizeUser(ctx, "DEL", tempUrl)
			if isAuthorized {
				ctx.JSON(http.StatusOK, _controller.DeleteABook(ctx, bookId))
			}
			tempUrl = hostURL
		})
	}

	// GET Clean up
	server.GET("/cleanup", func(ctx *gin.Context) {
		userId = 1
		bookId = 1
		tempUrl = hostURL
		userService = service.NewUserList()
		bookService = service.NewBookList()
		_controller = controller.NewController(userService, bookService)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "All are cleaned up!",
		})
	})

	server.Run()
}
