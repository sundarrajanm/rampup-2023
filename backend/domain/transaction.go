package domain

import "backend/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"tx_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"tx_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		Id:         t.TransactionId,
		Type:       t.TransactionType,
		Date:       t.TransactionDate,
		AccountId:  t.AccountId,
		NewBalance: t.Amount,
	}
}
