package controller

import (
	"net/http"

	"Desktop/lavina_tech_phase2/entity"
	"Desktop/lavina_tech_phase2/service"

	"github.com/gin-gonic/gin"
)

const messageFor401 string = "Unauthorized, permission denied"

type Controller interface {
	CreateNewUser(ctx *gin.Context, userId *int) gin.H
	AuthorizeUser(ctx *gin.Context, method string, url string) (bool, string)
	GetUserInfo(ctx *gin.Context) gin.H
	CreateABook(ctx *gin.Context, bookId *int, body string) gin.H
	GetAllBooks(ctx *gin.Context) gin.H
	EditABook(ctx *gin.Context, bookId string, body string) gin.H
	DeleteABook(ctx *gin.Context, bookId string) gin.H
}

type _controller struct {
	userService service.UserService
	bookService service.BookService
}

func NewController(userService service.UserService, bookService service.BookService) Controller {
	return &_controller{
		userService: userService,
		bookService: bookService,
	}
}

func (c *_controller) CreateNewUser(ctx *gin.Context, userId *int) gin.H {
	// Adding the user to the list of users
	var user entity.User = c.userService.AddUser(ctx, userId)

	return gin.H{
		"data":    user,
		"isOk":    true,
		"message": "ok",
	}
}

func (c *_controller) AuthorizeUser(ctx *gin.Context, method string, url string) (bool, string) {
	
	var isAuthorized bool
	var hash string
	var body string

	// Getting the user's "Secret" field
	var user entity.User = c.userService.FindUser(ctx)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, messageFor401)
		isAuthorized = false
	}
	var secret string = user.Secret

	// Checking whether the header "Sign" matches the hash
	hash, body = c.userService.GenerateMD5(ctx, method, url, secret)
	var sign string = ctx.Request.Header["Sign"][0]
	if hash == sign {
		isAuthorized = true
	} else {
		ctx.JSON(http.StatusUnauthorized, messageFor401)
		isAuthorized = false
	}
	return isAuthorized, body
}

func (c *_controller) GetUserInfo(ctx *gin.Context) gin.H {

	// Getting the user's information
	var user entity.User = c.userService.FindUser(ctx)

	return gin.H{
		"data":    user,
		"isOk":    true,
		"message": "ok",
	}
}

func (c *_controller) CreateABook(ctx *gin.Context, bookId *int, body string) gin.H {

	// Adding a new book to the bookshelf
	var bookInfo entity.Book = c.bookService.AddBook(ctx, bookId, body)

	return gin.H{
		"data":    bookInfo,
		"isOk":    true,
		"message": "ok",
	}
}

func (c *_controller) GetAllBooks(ctx *gin.Context) gin.H {

	// Getting the list of all books
	var books []entity.Book = c.bookService.ListBooks(ctx)

	return gin.H{
		"data":    books,
		"isOk":    true,
		"message": "ok",
	}
}

func (c *_controller) EditABook(ctx *gin.Context, bookId string, body string) gin.H {

	// Updating the book's status
	var bookInfo entity.Book = c.bookService.UpdateBook(ctx, bookId, body)

	return gin.H{
		"data":    bookInfo,
		"isOk":    true,
		"message": "ok",
	}
}

func (c *_controller) DeleteABook(ctx *gin.Context, bookId string) gin.H {

	// Removing the book from the bookshelf
	if c.bookService.RemoveBook(ctx, bookId) {
		return gin.H{
			"data":    "Successfully deleted",
			"isOk":    true,
			"message": "ok",
		}
	} else {
		return gin.H{
			"data":    "There is no book with ID: " + bookId,
			"isOk":    true,
			"message": "ok",
		}
	}
}
