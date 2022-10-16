package service

import (
	"github.com/dro14/lavina_tech_phase2/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookService interface {
	AddBook(ctx *gin.Context, bookId int) entity.BookInfo
	ListBooks() []entity.BookInfo
	UpdateBook(ctx *gin.Context, id string) entity.BookInfo
	RemoveBook(ctx *gin.Context, id string) bool
}

type bookService struct {
	books []entity.BookInfo
}

func NewBookList() BookService {
	return &bookService{}
}

func (service *bookService) AddBook(ctx *gin.Context, bookId int) entity.BookInfo {
	var tempMap = make(map[string]string)
	ctx.BindJSON(&tempMap)

	var isbn string = tempMap["isbn"]
	var url string = "https://openlibrary.org/isbn/" + isbn + ".json"

	response, _ := http.Get(url)
	body, _ := io.ReadAll(response.Body)
	var jsonData string = string(body)
	var book = make(map[string]interface{})
	json.Unmarshal([]byte(jsonData), &book)

	var bookInfo entity.BookInfo

	bookInfo.Book.ID = bookId
	isbn = fmt.Sprintf("%s", book["isbn_13"])
	bookInfo.Book.ISBN = isbn[1:len(isbn)-1]
	bookInfo.Book.Title = fmt.Sprintf("%s", book["title"])
	bookInfo.Book.Author = fmt.Sprintf("%s", book["authors"])
	published := fmt.Sprintf("%s", book["publish_date"])
	bookInfo.Book.Published, _ = strconv.Atoi(published[len(published)-4:])
	pages := fmt.Sprintf("%.0f", book["number_of_pages"])
	bookInfo.Book.Pages, _ = strconv.Atoi(pages)
	bookInfo.Status = 0

	service.books = append(service.books, bookInfo)

	return bookInfo
}

func (service *bookService) ListBooks() []entity.BookInfo {
	return service.books
}

func (service *bookService) UpdateBook(ctx *gin.Context, id string) entity.BookInfo {
	bookId, _ := strconv.Atoi(id)	
	var tempMap = make(map[string]int)
	ctx.BindJSON(&tempMap)
	for i, bookInfo := range service.books {
		if bookInfo.Book.ID ==  bookId {
			service.books[i].Status = tempMap["status"]
			return service.books[i]
		}
	}
	return entity.BookInfo{}
}

func (service *bookService) RemoveBook(ctx *gin.Context, id string) bool {
	bookId, _ := strconv.Atoi(id)
	var tempArr = []entity.BookInfo{}
	copy(tempArr, service.books)
	for i, bookInfo := range service.books {
		if bookInfo.Book.ID == bookId {
			copy(service.books[i:], service.books[i+1:])
			service.books[len(service.books)-1] = entity.BookInfo{}
			service.books = service.books[:len(service.books)-1]
			return true
		}
	}
	return false
}