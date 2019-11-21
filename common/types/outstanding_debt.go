package types

type OutstandingDebt string

type OutstandingDebtWithHeight struct {
	Height int64             `json:"height"`
	Result []OutstandingDebt `json:"result"`
}
