package models

type User struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
	Address string `json:"address"`
}
