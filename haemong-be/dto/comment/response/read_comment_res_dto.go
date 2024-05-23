package response

type ReadCommentResDTO struct {
	CommentId string `json:"commentId"`
	Comment   string `json:"comment"`
	UserId    string `json:"userId"`
	Date      string `json:"date"`
}
