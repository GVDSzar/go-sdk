package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgIssueDecreaseApproval to allow a registered owner
type MsgIssueDecreaseApproval struct {
	IssueId     string         `json:"issue_id" yaml:"issue_id"`
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      string        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueDecreaseApproval) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueDecreaseApproval) Type() string { return "issue_decrease_approval" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueDecreaseApproval) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("issueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.FromAddress.Equals(msg.ToAddress) {
		return sdk.ErrInvalidCoins("can't approve yourself")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueDecreaseApproval) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueDecreaseApproval) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}
