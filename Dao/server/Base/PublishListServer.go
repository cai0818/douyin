package Base

import (
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
)

func PublishList(user_id int) []*Response.VideoList {
	var db = model.DB
	list := []*Response.VideoList{}
	var Videos []vo.VideoList
	db.Where("author_id = ?", user_id).Find(&Videos)
	for i, video := range Videos {

		videoList := new(Response.VideoList) // 创建新的 VideoList 对象
		videoList.ID = i + 1
		videoList.Title = video.Title
		videoList.PlayURL = video.PlayURL
		videoList.CoverURL = video.CoverURL
		videoList.CommentCount = video.CommentCount
		var author vo.Author
		db.Find(&author, video.AuthorID)
		videoList.Author.ID = author.ID
		videoList.Author.Name = author.Name
		videoList.Author.BackgroundImage = author.BackgroundImage
		videoList.Author.FollowerCount = author.FollowerCount
		videoList.Author.FollowCount = author.FavoriteCount
		videoList.Author.Avatar = author.Avatar
		videoList.Author.Signature = author.Signature
		videoList.Author.TotalFavorited = author.TotalFavorited
		videoList.Author.WorkCount = author.WorkCount
		videoList.Author.FavoriteCount = author.FollowerCount
		var count int64
		db.Model(&vo.Favorite{}).Where("user_id = ? AND video_id = ?", author.ID, video.ID).Count(&count)
		if count > 0 {
			// 用户ID和视频ID的组合存在

			videoList.IsFavorite = true
		} else {
			// 用户ID和视频ID的组合不存在

			videoList.IsFavorite = false
		}
		//db.Model(&vo.IsFollower{}).Where("user_id = ? AND FollowerID = ?", author.ID, video.ID).Count(&count)
		//if count > 0 {
		//	// 用户ID和视频ID的组合存在
		//	fmt.Println("点赞")
		//	videoList.IsFavorite = true
		//} else {
		//	// 用户ID和视频ID的组合不存在
		//	fmt.Println("不点赞")
		//	videoList.IsFavorite = false
		//}
		list = append(list, videoList) // 将 videoList 添加到切片中
	}
	return list
}
