package model

type UserReq struct {
	TeamName string `json:"teamname"`
	Password string `json:"password"`
	Location string `json:"location"`
}

type UserResp struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type User struct {
	TeamName string
	Password string
	Location string
}

type CodeReq struct {
	Code string `json:"code"`
}
