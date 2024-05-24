package request

type CreateCommentReqDTO struct {
	Comment string `json:"comment"`
	UserId  string `json:"userId"`
	FeedId  string `json:"feedId"`
}
