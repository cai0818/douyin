package Base

import (
	"github.com/patrickmn/go-cache"
	"simpledouyin/Dao/Cache"
	"simpledouyin/Dao/vo"
	"simpledouyin/Response"
	"simpledouyin/model"
)

func FeedVideo() []*Response.VideoList {
	Cache.VideoMutex.Lock()
	if Cache.VideoList != nil { // 检查内存缓存
		copiedVideoList := make([]*Response.VideoList, len(Cache.VideoList))
		copy(copiedVideoList, Cache.VideoList)
		Cache.VideoMutex.Unlock()
		return copiedVideoList
	}
	Cache.VideoMutex.Unlock()

	// 尝试从缓存中获取数据
	if cachedList, found := Cache.C.Get("feed_video"); found {
		return cachedList.([]*Response.VideoList)
	}

	var db = model.DB
	list := []*Response.VideoList{}

	var Videos []vo.VideoList
	db.Order("id desc").Limit(30).Find(&Videos)
	for _, video := range Videos {
		videoList := new(Response.VideoList) // 创建新的 VideoList 对象
		videoList.ID = video.ID
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
			videoList.IsFavorite = true
		} else {
			videoList.IsFavorite = false
		}
		list = append(list, videoList) // 将 videoList 添加到切片中
	}

	Cache.C.Set("feed_video", list, cache.DefaultExpiration) // 将结果存入缓存

	Cache.VideoMutex.Lock()
	Cache.VideoList = make([]*Response.VideoList, len(list))
	copy(Cache.VideoList, list)
	Cache.VideoMutex.Unlock()

	return list
}
