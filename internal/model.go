package internal

type BaseResponse struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

type GetUnreadMessagesResponse struct {
	Data []UnreadMessage `json:"data"`
	BaseResponse
}

type UnreadMessage struct {
	From  string   `json:"from"`
	Texts []string `json:"text"`
}

type GetUserResponse struct {
	Data   []string `json:"data"`
	Status string   `json:"status"`
}

// DAO
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	LoggedIn bool   `json:"loggedin"`
	Password string `json:"password"`
}
