package src

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func VideoPost(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		log.Println("File upload error: ", err)
		c.JSON(http.StatusOK, RespHead{
			StatusCode: 1,
			StatusMsg:  "File upload error",
		})
		return
	}
	PostTime := time.Now().Unix()                                         //上传时间戳
	Route := "./Video/" + strconv.FormatInt(PostTime, 10) + file.Filename //存储路径
	err = c.SaveUploadedFile(file, Route)                                 //保存上传文件
	if err != nil {
		log.Println("File upload error: ", err)
		c.JSON(http.StatusOK, RespHead{
			StatusCode: 1,
			StatusMsg:  "File upload error",
		})
		return
	}

	Token := c.PostForm("token") //用户鉴权待完善
	Title := c.PostForm("title") //视频信息存储待完善

	c.JSON(http.StatusOK, RespHead{
		StatusCode: 0,
		StatusMsg:  "Succeed!",
	})
}
