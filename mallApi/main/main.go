package main

import (
	"github.com/gin-gonic/gin"
	"golangPractise/mallApi/controller"
	"net/http"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.GET("/hello", hello)

	r.POST("/user", controller.InsertNewUser)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
	})
}
