package Base

import "C"
import (
	"simpledouyin/Dao/Cache"
	"simpledouyin/Dao/vo"
	"simpledouyin/model"
	"time"
)

func UploadFileServer(filename string, userid int, title string, picturename string) {
	var db = model.DB
	videoList := new(vo.VideoList)
	videoList.CommentCount = 0

	videoList.PlayURL = "http://" + model.Arg + ":8080/static/" + filename
	videoList.CoverURL = "http://" + model.Arg + ":8080/static/picture/" + picturename + ".png"
	videoList.Title = title
	videoList.FavoriteCount = 0
	videoList.AuthorID = userid
	var author vo.Author
	db.Find(&author, userid)
	author.WorkCount++
	videoList.CreateTime = time.Now().Unix()
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
	db.Create(&videoList)
	db.Save(&author)
	Cache.VideoMutex.Lock()
	defer Cache.VideoMutex.Unlock()

	Cache.C.Delete("VideoList")
}
