package main

import (
	Config "bookmark/config"
	Mysql "bookmark/config/mysql"

	"github.com/gin-gonic/gin"
)

func main() {

	serverConfig := Config.LoadConfig()
	Mysql.InitMysql(serverConfig)

	r := gin.Default()

	r.LoadHTMLGlob("static/*.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
