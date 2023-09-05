package Base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Response"
	"simpledouyin/Tool"
	"strconv"
)

type Video struct {
	StatusCode int `json:"status_code"`
	//StatusMsg  string      `json:"status_msg"`
	VideoList []Response.VideoList `json:"video_list"`
}

func PublishList(c *gin.Context) {

	usrIdstr := c.Query("user_id")
	token := c.Query("token")
	code := Tool.Code(token)
	if code != 0 {
		fmt.Println(Tool.Code(token))
	}
	usrId, err := strconv.Atoi(usrIdstr)
	if err != nil {
		// 处理转换错误的情况
		// 例如返回错误响应或设置默认值
		// ...
	}
	result := new(Video)
	result.StatusCode = 0

	list := Base.PublishList(usrId)
	for _, item := range list {
		result.VideoList = append(result.VideoList, *item)
	}
	c.JSON(http.StatusOK, result)
}
