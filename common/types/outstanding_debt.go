package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OutstandingDebt sdk.Int

type OutstandingDebtWithHeight struct {
	Height int64             `json:"height"`
	Result []OutstandingDebt `json:"result"`
}
