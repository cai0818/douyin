package Interaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Interaction"
	"simpledouyin/Tool"
	"strconv"
)

func Like(c *gin.Context) {

	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	UserId := Tool.Code(token)
	id, err := strconv.Atoi(video_id)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	Interaction.Like(id, UserId, action_type)
	data := gin.H{"status_code": 0}

	c.JSON(http.StatusOK, data)
}
