package entity

type BookInfo struct {
	Book struct {
		ID        int    `json:"id"`
		ISBN      string `json:"isbn"`
		Title     string `json:"title"`
		Author    string `json:"author"`
		Published int    `json:"published"`
		Pages     int    `json:"pages"`
	} `json:"book"`
	Status int `json:"status"`
}