package entity

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type Book struct {
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

type BookInfo struct {
	ISBN_13 []string `json:"isbn_13"`
	Title   string `json:"title"`
	Authors []struct{Key string} `json:"authors"`
	Publish_date string `json:"publish_date"`
	Number_of_pages float64 `json:"number_of_pages"`
}

type ISBN struct {
	ISBN string `json:"isbn"`
}

type Status struct {
	Status int `json:"status"`
}

type Author struct {
	Name string `json:"name"`
}
