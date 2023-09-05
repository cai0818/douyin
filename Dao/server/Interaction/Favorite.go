package Interaction

import (
	"fmt"
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
)

func FavoriteList(user_id string) []*Response.VideoList {
	var db = model.DB

	var list []*Response.VideoList
	var favorites []vo.Favorite
	db.Where("user_id = ?", user_id).Find(&favorites)
	for _, favorite := range favorites {
		var videoList vo.VideoList
		tmp := new(Response.VideoList)
		db.Find(&videoList, favorite.VideoID)
		fmt.Println(videoList.PlayURL)
		tmp.ID = videoList.ID
		tmp.Title = videoList.Title
		tmp.PlayURL = videoList.PlayURL
		tmp.CoverURL = videoList.CoverURL
		tmp.CommentCount = videoList.CommentCount
		var author vo.Author
		db.Find(&author, videoList.AuthorID)
		tmp.Author.ID = author.ID
		tmp.Author.Name = author.Name
		tmp.Author.BackgroundImage = author.BackgroundImage
		tmp.Author.FollowerCount = author.FollowerCount
		tmp.Author.FollowCount = author.FavoriteCount
		tmp.Author.Avatar = author.Avatar
		tmp.Author.Signature = author.Signature
		tmp.Author.TotalFavorited = author.TotalFavorited
		tmp.Author.WorkCount = author.WorkCount
		tmp.Author.FavoriteCount = author.FollowerCount

		videoList.IsFavorite = true
		fmt.Println(tmp)
		list = append(list, tmp)
	}
	return list
}

func Like(VideoId int, UserID int, actionType string) {
	var db = model.DB
	var favorites vo.Favorite
	favorites.UserID = UserID
	favorites.VideoID = VideoId
	if actionType == "1" {
		db.Create(&favorites)
		addLike := new(vo.VideoList)
		db.Find(&addLike, VideoId)
		addLike.FavoriteCount++
		db.Save(&addLike)
	} else if actionType == "2" {
		db.Where("user_id = ? AND video_id = ?", UserID, VideoId).Delete(&favorites)
		addLike := new(vo.VideoList)
		db.Find(&addLike, VideoId)
		addLike.FavoriteCount--
		db.Save(&addLike)
	}

}
