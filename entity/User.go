package entity

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}