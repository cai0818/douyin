package Response

type CommentResult struct {
	StatusCode int     `json:"status_code"`
	Comment    Comment `json:"comment"`
}

type Comment struct {
	ID         int    `json:"id"`
	Author     User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentResultList struct {
	StatusCode  int           `json:"status_code"`
	CommentList []CommentList `json:"comment_list"`
}

type CommentList struct {
	ID         int    `json:"id"`
	Author     User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}
