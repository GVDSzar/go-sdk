package msg

import "github.com/cosmos/cosmos-sdk/types"

// MsgIssueIncreaseApproval to allow a registered owner
type IssueIncreaseApproval struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Int        `json:"amount" yaml:"amount"`
}
