package main

import (
	"Desktop/lavina_tech_phase2/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.RespondToSignup(ctx))
	})

	server.Run()
}