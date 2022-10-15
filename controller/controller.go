package controller

import (
	"Desktop/lavina_tech_phase2/entity"

	"github.com/gin-gonic/gin"
)

func RespondToSignup(ctx *gin.Context) gin.H {
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