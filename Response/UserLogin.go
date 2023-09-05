package Response

type UserLogin struct {
	StatusCode int    `json:"status_code"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}
