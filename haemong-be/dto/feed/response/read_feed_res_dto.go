package response

type ReadFeedResDTO struct {
	FeedId  string `json:"feedId"`
	UserId  string `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
