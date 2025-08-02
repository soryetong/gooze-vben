package dto

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Id          int64    `json:"id"`
	Password    string   `json:"password"`
	RealName    string   `json:"realName"`
	Roles       []string `json:"roles"`
	Username    string   `json:"username"`
	AccessToken string   `json:"accessToken"`
}

type LogoutReq struct {
	WithCredentials bool `json:"withCredentials"`
}
