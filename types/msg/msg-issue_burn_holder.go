package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgIssueBurnHolder struct {
	IssueId     string         `json:"issue_id" yaml:"issue_id"`
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Amount      string         `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueBurnHolder) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnHolder) Type() string { return "issue_burn_holder" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnHolder) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("issueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurnHolder) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnHolder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}
