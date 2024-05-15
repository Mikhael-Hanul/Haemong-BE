package request

type CreateFeedReqDTO struct {
	UserId  string `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
