package Interaction

import (
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
	"time"
)

func CommentAdd(UserId int, VideoId int, CommentText string) Response.Comment {
	var db = model.DB
	comment := new(vo.Comment)
	comment.UserID = UserId
	comment.VideoID = VideoId
	comment.Message = CommentText
	comment.CreateDate = time.Now()
	db.Create(&comment)
	//视频列表增加评论数量
	commentCount := new(vo.VideoList)
	db.Order("id desc").Find(&commentCount, VideoId)
	commentCount.CommentCount++
	db.Save(&commentCount)
	result := Response.Comment{
		ID:         comment.ID,
		Author:     Base.UserInfo(UserId),
		CreateDate: comment.CreateDate.Format("2006-01-02 15:04:05"),
		Content:    comment.Message,
	}

	return result
}
func CommentDele(UserId int, VideoId int, CommentId int) {
	var db = model.DB
	comment := vo.Comment{
		UserID:  UserId,
		VideoID: VideoId,
		ID:      CommentId,
	}
	db.Where(&comment).Delete(&comment)
	commentCount := new(vo.VideoList)
	db.Find(&commentCount, VideoId)
	commentCount.CommentCount--
	db.Save(&commentCount)
}

func CommentList(VideoId int) []*Response.CommentList {
	var db = model.DB
	var comments []vo.Comment
	db.Order("id desc").Where("VideoID = ?", VideoId).Find(&comments)
	list := []*Response.CommentList{}
	for _, comment := range comments {
		List := new(Response.CommentList)
		List.ID = comment.ID
		List.Author = Base.UserInfo(comment.UserID)
		List.Content = comment.Message
		List.CreateDate = comment.CreateDate.Format("2006-01-02 15:04:05")
		list = append(list, List)
	}
	return list
}
