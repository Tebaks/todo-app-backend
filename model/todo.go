package model

type Todo struct {
	ID   string `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}
