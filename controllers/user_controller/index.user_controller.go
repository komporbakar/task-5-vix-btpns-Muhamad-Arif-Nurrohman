package user_controller

import "github.com/gin-gonic/gin"

func GetUser(ctx *gin.Context) {
	isValidated := true

	if !isValidated {
		ctx.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"hello": "world kacang",
	})
}