package main

import (
	"awesomeProject/web-client/user_portal"
	"github.com/gin-gonic/gin"
)

func main() {

	user := user_portal.NewUserPortal()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {


		s := user.UserLogin()

		c.JSON(200, gin.H{
			"message": s,
		})
	})
	r.Run()
}
