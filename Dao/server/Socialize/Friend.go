package Socialize

import (
	"simpledouyin/Dao/server/Base"
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
)

func FriendList(UserId int) []*Response.UserList {
	var db = model.DB
	var followers []vo.IsFollower
	db.Model(&vo.IsFollower{}).Select("is_followers.user_id, is_followers.FollowerID").
		Joins("INNER JOIN is_followers AS uf ON is_followers.user_id = uf.FollowerID AND is_followers.FollowerID = uf.user_id").
		Find(&followers)
	list := []*Response.UserList{}
	for _, follower := range followers {
		if follower.UserID == UserId {
			continue
		}
		userInfo := Base.UserInfo(follower.UserID)
		tmp := new(Response.UserList)
		tmp.ID = userInfo.ID
		tmp.Name = userInfo.Name
		tmp.FollowCount = userInfo.FollowerCount
		userInfo.IsFollow = true
		tmp.Avatar = userInfo.Avatar
		tmp.Signature = userInfo.Signature
		tmp.TotalFavorited = userInfo.TotalFavorited
		tmp.WorkCount = userInfo.WorkCount
		tmp.FavoriteCount = userInfo.FavoriteCount
		tmp.BackgroundImage = userInfo.BackgroundImage
		list = append(list, tmp)
	}
	return list
}
