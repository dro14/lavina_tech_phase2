package controller

import (
	"Desktop/lavina_tech_phase2/entity"

	"github.com/gin-gonic/gin"
)

func CreateNewUser(ctx *gin.Context) gin.H {
	var id = 32
	var userData entity.User
	ctx.BindJSON(&userData)

	return gin.H{
		"data": gin.H{
			"id": id,
			"name": userData.Name,
			"key": userData.Key,
			"secret": userData.Secret,
		},
		"isOk": true,
		"message": "ok",
	}
}

func GetUserInfo(ctx *gin.Context) gin.H {
	var id = 32
	var userData entity.User
	ctx.BindJSON(&userData)

	return gin.H{
		"data": gin.H{
			"id": id,
			"name": userData.Name,
			"key": userData.Key,
			"secret": userData.Secret,
		},
		"isOk": true,
		"message": "ok",
	}
}

func CreateABook(ctx *gin.Context) gin.H {
	var id = 32
	var userData entity.User
	ctx.BindJSON(&userData)

	return gin.H{
		"data": gin.H{
			"id": id,
			"name": userData.Name,
			"key": userData.Key,
			"secret": userData.Secret,
		},
		"isOk": true,
		"message": "ok",
	}
}

func GetAllBooks(ctx *gin.Context) gin.H {
	var id = 32
	var userData entity.User
	ctx.BindJSON(&userData)

	return gin.H{
		"data": gin.H{
			"id": id,
			"name": userData.Name,
			"key": userData.Key,
			"secret": userData.Secret,
		},
		"isOk": true,
		"message": "ok",
	}
}

func EditABook(ctx *gin.Context) gin.H {
	var id = 32
	var userData entity.User
	ctx.BindJSON(&userData)

	return gin.H{
		"data": gin.H{
			"id": id,
			"name": userData.Name,
			"key": userData.Key,
			"secret": userData.Secret,
		},
		"isOk": true,
		"message": "ok",
	}
}

func DeleteABook(ctx *gin.Context) gin.H {
	var id = 32
	var userData entity.User
	ctx.BindJSON(&userData)

	return gin.H{
		"data": gin.H{
			"id": id,
			"name": userData.Name,
			"key": userData.Key,
			"secret": userData.Secret,
		},
		"isOk": true,
		"message": "ok",
	}
}