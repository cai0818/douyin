package Base

import (
	"fmt"
	"simpledouyin/Dao/vo"
	"simpledouyin/model"
)

func Login(username string, password string) *vo.Author {
	var db = model.DB
	var author vo.Author

	// 执行查询操作
	result := db.Where("name = ? and password = ?", username, password).First(&author)

	if result.Error != nil {
		// 处理查询错误
		// 可以返回错误信息给客户端或执行其他操作
		fmt.Println(result.Error)
		return nil
	} else if result.RowsAffected == 0 {
		// 处理查询结果为空
		return nil
	} else {
		// 查询成功，返回查询结果
		return &author
	}
}
