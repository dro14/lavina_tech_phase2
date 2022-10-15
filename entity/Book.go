package entity

type Book struct {
	ID        int    `json:"id"`
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published int    `json:"published"`
	Pages     int    `json:"pages"`
}