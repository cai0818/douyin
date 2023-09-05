package vo

type IsFollower struct {
	ID       int
	UserID   int `gorm:"column:user_id"`    //用户
	FollerID int `gorm:"column:FollowerID"` //被关注人id
}
