package Socialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Socialize"
	"simpledouyin/Response"
	"simpledouyin/Tool"
	"strconv"
)

func MessageList(c *gin.Context) {
	token := c.Query("token")
	ToUserIDStr := c.Query("to_user_id")
	UserId := Tool.Code(token)
	ToUserID, _ := strconv.Atoi(ToUserIDStr)
	fmt.Print(UserId, ToUserID)
	list := Socialize.MessageList(UserId, ToUserID)
	result := new(Response.Message)
	result.StatusCode = "0"
	result.MessageList = []Response.MessageList{}

	for _, item := range list {
		result.MessageList = append(result.MessageList, *item)
	}
	c.JSON(http.StatusOK, result)
}
func AddMessage(c *gin.Context) {
	token := c.Query("token")
	ToUserIDStr := c.Query("to_user_id")
	UserId := Tool.Code(token)
	ToUserID, _ := strconv.Atoi(ToUserIDStr)
	content := c.Query("content")
	Socialize.AddMessage(UserId, ToUserID, content)
	data := gin.H{"status_code": 0}
	c.JSON(http.StatusOK, data)
}
