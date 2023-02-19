package dto

import "backend/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	Type       string  `json:"transaction_type"`
	Amount     float64 `json:"amount"`
	CustomerId string
	AccountId  string
}

func (t TransactionRequest) IsWithdrawal() bool {
	return t.Type == WITHDRAWAL
}

func (r TransactionRequest) Validate() *errs.AppError {

	if r.Type != WITHDRAWAL && r.Type != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil
}
