package request

type UserReqDTO struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
