package entity

type User struct {
	ID     int
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}