package response

type ReadFeedResDTO struct {
	FeedId       string `json:"feedId"`
	UserId       string `json:"userId"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	LikeCount    int    `json:"likeCount"`
	DislikeCount int    `json:"dislikeCount"`
}
