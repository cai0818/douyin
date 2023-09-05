package Response

type Follower struct {
	StatusCode string   `json:"status_code"`
	UserList   []Author `json:"user_list"`
}
