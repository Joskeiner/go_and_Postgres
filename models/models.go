package models

type Stock struct {
	StockID int64  `json:"strockid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}
