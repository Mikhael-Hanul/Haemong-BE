package response

type SearchIlgiResDTO struct {
	IlgiId  string `json:"ilgiId"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Weather string `json:"weather"`
}
