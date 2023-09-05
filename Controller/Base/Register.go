package Base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Response"
	"simpledouyin/Tool"
)

type Result struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var Author = Base.Register(username, password)

	if Author != nil {

		token := Tool.Encode(Author.ID)
		result := new(Response.UserLogin)
		result.StatusCode = 0
		result.UserId = Author.ID
		result.Token = token
		c.JSON(http.StatusOK, result)
	} else {
		data := gin.H{"status_code": 0,
			"status_msg": "已经存在用户"}

		c.JSON(http.StatusOK, data)
	}
}
