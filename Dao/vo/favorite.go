package vo

type Favorite struct {
	ID      int
	UserID  int `gorm:"column:user_id"`
	VideoID int `gorm:"column:video_id"`
}
