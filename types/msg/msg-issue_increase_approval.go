package msg

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

// MsgIssueIncreaseApproval to allow a registered owner
type MsgIssueIncreaseApproval struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Int        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueIncreaseApproval) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueIncreaseApproval) Type() string { return "issue_increase_approval" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueIncreaseApproval) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("issueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.Amount.IsNegative() {
		return errors.New("can't approve negative coin amount")
	}
	if msg.FromAddress.Equals(msg.ToAddress) {
		return errors.New("can't approve yourself")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueIncreaseApproval) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueIncreaseApproval) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
