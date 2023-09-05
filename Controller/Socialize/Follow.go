package Socialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Socialize"
	"simpledouyin/Response"
	"simpledouyin/Tool"
	"strconv"
)

func Follow(c *gin.Context) {
	token := c.Query("token")
	ToUserIDStr := c.Query("to_user_id")
	ActionType := c.Query("action_type")
	UserId := Tool.Code(token)
	ToUserID, _ := strconv.Atoi(ToUserIDStr)
	Socialize.DoFollow(UserId, ToUserID, ActionType)

	data := gin.H{"status_code": 0}
	c.JSON(http.StatusOK, data)
}
func FollowList(c *gin.Context) {
	UserIdStr := c.Query("user_id")
	UserID, _ := strconv.Atoi(UserIdStr)
	followList := Socialize.FollowList(UserID)
	result := new(Response.Follower)
	result.StatusCode = "0"
	result.UserList = []Response.Author{}

	for _, item := range followList {
		result.UserList = append(result.UserList, *item)
	}
	// 将结果转换为JSON并发送响应
	c.JSON(http.StatusOK, result)
}
