package schema

type LoginReq struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResp struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
