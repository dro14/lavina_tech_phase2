package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/dro14/lavina_tech_phase2/entity"
	"github.com/dro14/lavina_tech_phase2/service"

	"github.com/gin-gonic/gin"
)

var userService service.UserService = service.NewUserList()

func Authorize(ctx *gin.Context, method string, url string) bool {
	var user entity.User = userService.FindUser(ctx)
	if user.ID == 0 {
		return false
	}

	request, _ := io.ReadAll(ctx.Request.Body)

	var body string = string(request)
	body = strings.ReplaceAll(body, " ", "")
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.ReplaceAll(body, "\t", "")

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