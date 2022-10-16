package controller

import (
	"Desktop/lavina_tech_phase2/entity"
	"Desktop/lavina_tech_phase2/service"

	"github.com/gin-gonic/gin"
)	

var bookService service.BookService = service.NewBookList()

func CreateNewUser(ctx *gin.Context, userId int) gin.H {
	var user entity.User = userService.AddUser(ctx, userId)
	return gin.H{
		"data": user,
		"isOk": true,
		"message": "ok",
	}
}

func GetUserInfo(ctx *gin.Context) gin.H {
	var user entity.User = userService.FindUser(ctx)
	return gin.H{
		"data": user,
		"isOk": true,
		"message": "ok",
	}
}

func CreateABook(ctx *gin.Context, bookId int) gin.H {
	var bookInfo entity.BookInfo = bookService.AddBook(ctx, bookId)
	return gin.H{
		"data": bookInfo,
		"isOk": true,
		"message": "ok",
	}
}

func GetAllBooks() gin.H{
	var books []entity.BookInfo = bookService.ListBooks()
	return gin.H{
		"data": books,
		"isOk": true,
		"message": "ok",
	}
}

func EditABook(ctx *gin.Context, bookId string) gin.H {
	var bookInfo entity.BookInfo = bookService.UpdateBook(ctx, bookId)
	return gin.H{
		"data": bookInfo,
		"isOk": true,
		"message": "ok",
	}
}

func DeleteABook(ctx *gin.Context, bookId string) gin.H {
	if bookService.RemoveBook(ctx, bookId) {
		return gin.H{
			"data": "Successfully deleted",
			"isOk": true,
			"message": "ok",
		} 
	} else {
		return gin.H{
			"data": "There is no book with such ID",
			"isOk": true,
			"message": "ok",
		}
	}
}