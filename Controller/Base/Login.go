package Base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Response"
	"simpledouyin/Tool"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	author := Base.Login(username, password)
	if author != nil {
		token := Tool.Encode(author.ID)
		result := new(Response.UserLogin)
		result.StatusCode = 0
		result.UserId = author.ID
		result.Token = token
		c.JSON(http.StatusOK, result)
	}
}
