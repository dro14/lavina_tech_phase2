package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"Desktop/lavina_tech_phase2/entity"

	"github.com/gin-gonic/gin"
)

type BookService interface {
	AddBook(ctx *gin.Context, bookId *int, body string) entity.Book
	ListBooks(ctx *gin.Context) []entity.Book
	UpdateBook(ctx *gin.Context, id string, body string) entity.Book
	RemoveBook(ctx *gin.Context, id string) bool
}

type bookService struct {
	books []entity.Book
}

func NewBookList() BookService {
	return &bookService{}
}

func (service *bookService) AddBook(ctx *gin.Context, bookId *int, body string) entity.Book {


	// Retrieving the book's ISBN from the body of the request
	var temp entity.ISBN
	err := json.Unmarshal([]byte(body), &temp)
	if err != nil {
		log.Fatalln(err)
	}
	var isbn string = temp.ISBN

	// if len(isbn) == 0 {
	// 	jsonData, _ := io.ReadAll(ctx.Request.Body)
	// 	log.Fatalln("Body: " + string(jsonData))
	// }

	// Sending "GET" OpenLibrary API request
	var url string = "https://openlibrary.org/isbn/" + isbn + ".json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// Converting the body of the request into the string format
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	
	// Adjusting the format of the response
	var bookInfo entity.BookInfo
	err = json.Unmarshal(jsonData, &bookInfo)
	if err != nil {
		// log.Fatalln("Hi")
		ctx.String(500, string(jsonData) + "Hello")

	}

	// Retriving the relevant data
	var book entity.Book

	// 1. Assigning an ID to the book
	book.Book.ID = *bookId
	// 2. Getting ISBN
	book.Book.ISBN = bookInfo.ISBN_13[0]
	// 3. Getting the title
	book.Book.Title = bookInfo.Title
	// 4. Getting the authors' names
	for i := range bookInfo.Authors {
		// Sending "GET" OpenLibrary API request
		url = "https://openlibrary.org" + bookInfo.Authors[i].Key + ".json"
		response, err = http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		// Converting the body of the request into the string format
		jsonData, err = io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		// Adjusting the format of the response
		var author entity.Author
		err = json.Unmarshal(jsonData, &author)
		if err != nil {
			log.Fatalln(err)
		}
		// Retriving the relevant data
		var name string = author.Name
		if len(name) > 0 {
			if i > 0 {
				book.Book.Author += ", "
			}
			book.Book.Author += name
		}
	}
	// 5. Getting the year of publication
	published := bookInfo.Publish_date
	if len(published) > 0 {
	 	book.Book.Published, err = strconv.Atoi(published[len(published)-4:])
		if err != nil {
			book.Book.Published = 0
		}
	}
	// 6. Getting the number of pages
	pages := fmt.Sprintf("%.0f", bookInfo.Number_of_pages)
	book.Book.Pages, err = strconv.Atoi(pages)
	if err != nil {
		book.Book.Pages = 0
	}
	// Assigning status of 0 to the book because it is new
	book.Status = 0

	// Checking for duplicates
	for _, value := range service.books {
		if value.Book.Title == bookInfo.Title {
			if value.Book.ISBN == bookInfo.ISBN_13[0] {
				book.Book.ID = 0
				break
			}
		}
	}

	// Adding the book to the bookshelf
	if book.Book.ID != 0 {
		service.books = append(service.books, book)
		*bookId++
	}
	return book
}

func (service *bookService) ListBooks(ctx *gin.Context) []entity.Book {
	return service.books
}

func (service *bookService) UpdateBook(ctx *gin.Context, id string, body string) entity.Book {

	// Retriving the book's ID from URL and converting to integer format
	bookId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	// Retrieving the status to be assigned to the book from the body of the request
	var temp entity.Status
	err = json.Unmarshal([]byte(body), &temp)
	if err != nil {
		log.Fatalln(err)
	}
	// Updating the status of the book
	var book entity.Book
	for i, value := range service.books {
		if value.Book.ID == bookId {
			service.books[i].Status = temp.Status
			return service.books[i]
		}
	}
	book.Book.ID = 0
	return book
}

func (service *bookService) RemoveBook(ctx *gin.Context, id string) bool {

	// Retriving the book's ID from URL and converting to integer format
	bookId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}
	// Removing the book from the bookshelf
	for i, bookInfo := range service.books {
		if bookInfo.Book.ID == bookId {
			copy(service.books[i:], service.books[i+1:])
			service.books[len(service.books)-1] = entity.Book{}
			service.books = service.books[:len(service.books)-1]
			return true
		}
	}
	return false
}
