package Base

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"path/filepath"
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Tool"
	"strings"
	"time"
)

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("data")
	token := c.PostForm("token")
	title := c.PostForm("title")
	data := gin.H{"status_code": 0}
	if err != nil {
		fmt.Printf("FormFile,err=%v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	userid := Tool.Code(token)
	fmt.Println(title)
	fmt.Println(userid)
	fmt.Println(file.Filename)

	// 生成保存路径和文件名
	storagePath := "static"                     // 设置存储目录的路径
	filename := generateFileName(file.Filename) // 生成唯一的文件名

	// 将文件保存到磁盘上的指定目录
	err = c.SaveUploadedFile(file, filepath.Join(storagePath, filename))
	if err != nil {
		fmt.Printf("SaveUploadedFile,err=%v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	snapshotPath := strings.Replace(filename, ".mp4", "", -1)
	videopath := storagePath + "/" + filename
	Tool.GetSnapshot(videopath, snapshotPath, 1)
	Base.UploadFileServer(filename, userid, title, snapshotPath)
	c.JSON(http.StatusOK, data)
}

// 生成唯一的文件名
func generateFileName(originalFilename string) string {
	// 使用时间戳和随机数来生成唯一的文件名
	timestamp := time.Now().Unix()
	randomNumber := rand.Intn(10000)
	filename := fmt.Sprintf("%d-%d-%s", timestamp, randomNumber, originalFilename)
	return filename
}
