package service

import (
	"github.com/dro14/lavina_tech_phase2/entity"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	AddUser(ctx *gin.Context, id int) entity.User
	FindUser(ctx *gin.Context) entity.User
}

type userService struct {
	users []entity.User
}

func NewUserList() UserService {
	return &userService{}
}

func (service *userService) AddUser(ctx *gin.Context, userId int) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	user.ID = userId
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindUser(ctx *gin.Context) entity.User {
	var key = ctx.Request.Header["Key"][0]
	for _ , user := range service.users {
		if user.Key == key {
			return user
		}
	}
	return entity.User{ ID: 0, }
}