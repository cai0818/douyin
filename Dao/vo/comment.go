package vo

import "time"

type Comment struct {
	ID         int
	UserID     int       `gorm:"column:user_id"` //用户
	VideoID    int       `gorm:"column:VideoID"` //评论视频ID
	Message    string    `gorm:"column:Message"`
	CreateDate time.Time `gorm:"column:CreateDate"`
}
