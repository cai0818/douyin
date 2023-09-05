package Socialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Socialize"
	"simpledouyin/Response"
	"strconv"
)

func FanList(c *gin.Context) {

	UserIdStr := c.Query("user_id")
	UserID, _ := strconv.Atoi(UserIdStr)
	followList := Socialize.FunList(UserID)
	result := new(Response.Follower)
	result.StatusCode = "0"
	result.UserList = []Response.Author{}

	for _, item := range followList {
		result.UserList = append(result.UserList, *item)
	}
	// 将结果转换为JSON并发送响应
	c.JSON(http.StatusOK, result)
}
