package controller

import (
	"github.com/dro14/lavina_tech_phase2/entity"
	"github.com/dro14/lavina_tech_phase2/service"
	"crypto/md5"
	"fmt"

	"github.com/gin-gonic/gin"
)

var userService service.UserService = service.NewUserList()

func Authorize(method string, url string) bool {
	var ctx *gin.Context
	
	var user entity.User = userService.FindUser(ctx)
	if user.ID == 0 {
		return false
	}

	var body string
	ctx.BindJSON(&body)
	var secret string = user.Secret
	
	var stringToSign string = method + url + body + secret
	var data = md5.Sum([]byte(stringToSign))
	var resultant_hash string
	for _, i := range data {
		resultant_hash += fmt.Sprintf("%x", i)
	}

	var sign string = ctx.Request.Header["Sign"][0]
	if resultant_hash == sign {
		return true
	} else {
		return false
	}
}