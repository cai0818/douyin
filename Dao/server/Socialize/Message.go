package Socialize

import (
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
	"time"
)

func MessageList(UserId int, ToUserId int) []*Response.MessageList {
	var db = model.DB
	var messages []vo.Message
	db.Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
		UserId, ToUserId, ToUserId, UserId).
		Find(&messages)
	list := []*Response.MessageList{}
	for _, message := range messages {
		tmp := new(Response.MessageList)
		tmp.ID = message.ID
		tmp.Content = message.Content
		tmp.CreateTime = message.CreateTime
		tmp.FromUserID = message.FromUserID
		tmp.ToUserID = message.ToUserID
		list = append(list, tmp)
	}
	return list
}
func AddMessage(UserId int, ToUserId int, content string) {
	var db = model.DB
	var messages vo.Message
	messages.FromUserID = UserId
	messages.ToUserID = ToUserId
	messages.Content = content
	messages.CreateTime = time.Now().Unix()
	db.Create(&messages)
}
