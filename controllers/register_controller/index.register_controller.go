package register_user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	isValidated := true

	if !isValidated {
		ctx.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	// ctx.HTML("suksesss")
	
		fmt.Println("Halaman  sukses dibuat")
	
}