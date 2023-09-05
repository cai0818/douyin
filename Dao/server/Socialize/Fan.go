package Socialize

import (
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
)

func FunList(UserId int) []*Response.Author {
	var followers []vo.IsFollower
	var db = model.DB
	list := []*Response.Author{}
	db.Where("FollowerID = ?", UserId).Find(&followers)
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
