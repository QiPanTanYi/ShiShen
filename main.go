package main

import (
	"shishen/src"
	"github.com/gin-gonic/gin"
)

func main() {
	main := gin.Default()
	main.POST("/douyin/publish/action/",src.VideoPost)
	main.Run()
}
