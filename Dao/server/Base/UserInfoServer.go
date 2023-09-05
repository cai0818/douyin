package Base

import (
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
)

func UserInfo(user_id int) Response.User {
	var db = model.DB
	var author vo.Author
	var user Response.User
	db.Find(&author, user_id)
	user.ID = author.ID
	user.Name = author.Name
	user.BackgroundImage = author.BackgroundImage
	user.FollowerCount = author.FollowerCount
	user.FollowCount = author.FavoriteCount
	user.Avatar = author.Avatar
	user.Signature = author.Signature
	user.TotalFavorited = author.TotalFavorited
	user.WorkCount = author.WorkCount
	user.FavoriteCount = author.FollowerCount
	return user
}
