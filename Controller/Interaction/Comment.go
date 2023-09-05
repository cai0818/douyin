package Interaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/Dao/server/Interaction"
	"simpledouyin/Response"
	"simpledouyin/Tool"
	"strconv"
)

func Conmment(c *gin.Context) {
	token := c.Query("token")
	VideoIdStr := c.Query("video_id")
	ActionType := c.Query("action_type")
	UserId := Tool.Code(token)
	VideoId, _ := strconv.Atoi(VideoIdStr)
	if ActionType == "1" {
		CommentText := c.Query("comment_text")
		comment := Interaction.CommentAdd(UserId, VideoId, CommentText)
		result := new(Response.CommentResult)
		result.Comment = comment
		result.StatusCode = 0
		c.JSON(http.StatusOK, result)
	} else if ActionType == "2" {
		CommentIdStr := c.Query("comment_id")
		CommentID, _ := strconv.Atoi(CommentIdStr)
		Interaction.CommentDele(UserId, VideoId, CommentID)
		commentList := Interaction.CommentList(VideoId)

		result := new(Response.CommentResultList)
		for _, item := range commentList {
			result.CommentList = append(result.CommentList, *item)
		}
		result.StatusCode = 0
		c.JSON(http.StatusOK, result)
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	UserId := Tool.Code(token)
	fmt.Println(UserId)

	VideoIdStr := c.Query("video_id")
	VideoId, _ := strconv.Atoi(VideoIdStr)

	commentList := Interaction.CommentList(VideoId)

	result := new(Response.CommentResultList)
	for _, item := range commentList {
		result.CommentList = append(result.CommentList, *item)
	}
	result.StatusCode = 0
	c.JSON(http.StatusOK, result)

}
