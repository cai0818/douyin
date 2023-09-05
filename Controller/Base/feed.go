package Base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Response"
	"time"
)

type test struct {
	StatusCode int         `json:"status_code"`
	VideoList  []VideoList `json:"video_list"`
	NextTime   int         `json:"next_time"`
}
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type VideoList struct {
	ID       int    `json:"id"`
	Author   Author `json:"author"`
	PlayURL  string `json:"play_url"`
	CoverURL string `json:"cover_url"`
}

func Feed(c *gin.Context) {
	//var db = model.DB

	result := new(Response.Video)
	result.StatusCode = 0

	result.NextTime = int(time.Now().Unix())

	list := Base.FeedVideo()
	for _, item := range list {
		result.VideoList = append(result.VideoList, *item)
	}
	c.JSON(http.StatusOK, result)

}
