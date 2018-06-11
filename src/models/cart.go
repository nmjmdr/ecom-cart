package models

type Item struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
}

type Cart struct {
	Id    string `json:"id"`
	Items []Item `json:"items"`
}
