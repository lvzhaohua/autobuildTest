package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.New()
	engine.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello gin ...")
	})
	engine.GET("/hello1", func(context *gin.Context) {
		context.String(200, "hello1 gin ...")
	})
	engine.Run(":8200")

}
