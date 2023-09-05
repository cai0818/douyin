package Response

type Message struct {
	StatusCode  string        `json:"status_code"`
	MessageList []MessageList `json:"message_list"`
}
type MessageList struct {
	ID         int    `json:"id"`
	ToUserID   int    `json:"to_user_id"`
	FromUserID int    `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
