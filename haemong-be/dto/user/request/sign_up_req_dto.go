package request

type SignUpReqDTO struct {
	UserId   string `json:"userId"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
