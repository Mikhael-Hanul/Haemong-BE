package request

type SaveIlgiReqDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Weather string `json:"weather"`
}
