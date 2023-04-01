package router

import (
	register_user "gin-gonic-gorm/controllers/register_controller"
	"gin-gonic-gorm/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {

	route := app

	route.GET("/", user_controller.GetUser)
	route.GET("/users/login", user_controller.GetUser)
	route.GET("/users/register", register_user.RegisterUser)
}