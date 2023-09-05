package Response

type Video struct {
	StatusCode int `json:"status_code"`
	//StatusMsg  string      `json:"status_msg"`
	NextTime  int         `json:"next_time"`
	VideoList []VideoList `json:"video_list"`
}
type Author struct {
	ID              int    `json:"id" `
	Name            string `json:"name"`
	Password        string `json:"-"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}

type VideoList struct {
	ID            int    `json:"id"`
	Author        Author `json:"author" `
	AuthorID      int    `json:"-"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
