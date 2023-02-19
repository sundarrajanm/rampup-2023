package dto

type TransactionResponse struct {
	Id         string  `json:"transaction_id"`
	NewBalance float64 `json:"new_balance"`
	Date       string  `json:"transaction_date"`
	AccountId  string  `json:"account_id"`
	Type       string  `json:"transaction_type"`
}
