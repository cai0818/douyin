package Socialize

import (
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
)

func DoFollow(UserId int, ToUserId int, actionType string) {
	var db = model.DB
	addFollow := new(vo.IsFollower)
	addFollow.UserID = UserId
	addFollow.FollerID = ToUserId
	if actionType == "1" {
		db.Create(&addFollow)
		Author := new(vo.Author)
		Author.ID = ToUserId
		db.Find(&Author)
		Author.FollowerCount++
		db.Save(&Author)
	} else if actionType == "2" {

		db.Where("user_id = ? AND FollowerID = ?", UserId, ToUserId).Delete(&addFollow)
	}
}
func FollowList(UserId int) []*Response.Author {

	var followers []vo.IsFollower
	var db = model.DB
	list := []*Response.Author{}
	db.Where("user_id = ?", UserId).Find(&followers)
	for _, follower := range followers {
		var Author vo.Author
		Author.ID = follower.FollerID
		db.Find(&Author, Author.ID)
		List := new(Response.Author)
		List.ID = Author.ID
		List.ID = Author.ID
		List.Name = Author.Name
		List.BackgroundImage = Author.BackgroundImage
		List.FollowerCount = Author.FollowerCount
		List.FollowCount = Author.FavoriteCount
		List.Avatar = Author.Avatar
		List.Signature = Author.Signature
		List.TotalFavorited = Author.TotalFavorited
		List.WorkCount = Author.WorkCount
		List.FavoriteCount = Author.FollowerCount
		List.IsFollow = true
		list = append(list, List)
	}
	return list

}
