package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"Desktop/lavina_tech_phase2/entity"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	AddUser(ctx *gin.Context, userId *int) entity.User
	FindUser(ctx *gin.Context) entity.User
	GenerateMD5(ctx *gin.Context, method string, url string, secret string) (string, string)
}

type userService struct {
	users []entity.User
}

func NewUserList() UserService {
	return &userService{}
}

func (service *userService) AddUser(ctx *gin.Context, userId *int) entity.User {

	// Retrieving the user's information from the body of the request
	var user entity.User
	ctx.BindJSON(&user)
	user.ID = *userId

	// Checking for duplicates
	var isIncluded bool = false
	for _, value := range service.users {
		if value.Name == user.Name {
			if value.Key == user.Key {
				if value.Secret == user.Secret {
					isIncluded = true
					break
				}
			}
		}
	}

	// Adding the user to the list of users
	if isIncluded {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "The user is already on the list",
		})
		return entity.User{}
	} else {
		service.users = append(service.users, user)
		*userId++
		return user
	}
}

func (service *userService) FindUser(ctx *gin.Context) entity.User {

	// Finding the user from the list based on its key
	var key string = ctx.Request.Header["Key"][0]

	for _, user := range service.users {
		if user.Key == key {
			return user
		}
	}

	return entity.User{ID: 0}
}

func (service *userService) GenerateMD5(ctx *gin.Context, method string, url string, secret string) (string, string) {

	// Extracting the body of the request and converting it into the string format
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var body string = string(jsonData)
	var trimmedBody string
	trimmedBody = strings.ReplaceAll(body, " ", "")
	trimmedBody = strings.ReplaceAll(trimmedBody, "\n", "")
	trimmedBody = strings.ReplaceAll(trimmedBody, "\t", "")
	trimmedBody = strings.ReplaceAll(trimmedBody, "{\"", "{")
	trimmedBody = strings.ReplaceAll(trimmedBody, ",\"", ",")
	trimmedBody = strings.ReplaceAll(trimmedBody, "\":", ":")

	// Encrypting the string to sign using MD5 algorithm
	var stringToSign string = method + url + body + secret
	var data = md5.Sum([]byte(stringToSign))
	var hash string
	for _, value := range data {
		hash += fmt.Sprintf("%x", value)
	}

	if len(body) == 0 {
		body = "{}"
	}

	return hash, body
}
