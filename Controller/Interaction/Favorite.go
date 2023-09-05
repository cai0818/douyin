package Interaction

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Interaction"
	"simpledouyin/Response"
	"time"
)

func Favorite(c *gin.Context) {

	userid := c.Query("user_id")
	result := new(Response.Video)
	result.StatusCode = 0

	result.NextTime = int(time.Now().Unix())
	list := Interaction.FavoriteList(userid)
	for _, item := range list {
		result.VideoList = append(result.VideoList, *item)
	}
	c.JSON(http.StatusOK, result)
}
