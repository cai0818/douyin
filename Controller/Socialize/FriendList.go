package Socialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Socialize"
	"simpledouyin/Response"
	"strconv"
)

func FriendList(c *gin.Context) {
	UserIdStr := c.Query("user_id")
	UserID, _ := strconv.Atoi(UserIdStr)
	list := Socialize.FriendList(UserID)

	result := new(Response.UserInfoList)
	result.StatusCode = "0"
	result.UserList = []Response.UserList{}

	for _, item := range list {
		result.UserList = append(result.UserList, *item)
	}
	c.JSON(http.StatusOK, result)
}
