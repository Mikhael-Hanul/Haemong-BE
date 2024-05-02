package request

type ChangePasswordReqDTO struct {
	UserId      string `json:"userId"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
