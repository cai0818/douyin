package Base

import (
	"simpledouyin/Dao/vo"
	"simpledouyin/model"
)

func Register(username string, password string) *vo.Author {
	var db = model.DB
	var author vo.Author

	// 检查数据库中是否已经存在相同username的记录
	result := db.Where("name = ?", username).Limit(1).Find(&author)

	if result.RowsAffected == 0 {
		// 在不存在相同username的情况下才进行插入操作
		author.Name = username
		author.Password = password

		author.Avatar = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
		author.TotalFavorited = "0"
		author.Signature = "该用户没用简介"
		author.BackgroundImage = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
		db.Create(&author)
		return &author
	}

	return nil
}
