package bootstrap

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/router"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	database.ConnectDatabase()
	app := gin.Default()
	router.InitRouter(app)
	app.Run(":3200")
}