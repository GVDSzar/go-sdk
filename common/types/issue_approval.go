package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Approval struct {
	Amount sdk.Int `json:"amount"`
}

type IssueApprovalWithHeight struct {
	Height int64      `json:"height"`
	Result []Approval `json:"result"`
}
