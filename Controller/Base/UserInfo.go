package Base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Response"
	"strconv"
)

func UserInfo(c *gin.Context) {
	usrIdstr := c.Query("user_id")
	//token := c.Query("token")
	usrId, err := strconv.Atoi(usrIdstr)
	if err != nil {
		// 处理转换错误的情况
		// 例如返回错误响应或设置默认值
		// ...
	}
	result := new(Response.UserInfo)
	result.StatusCode = 0
	result.User = Base.UserInfo(usrId)
	c.JSON(http.StatusOK, result)
}
