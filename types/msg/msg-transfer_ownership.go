package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const issueRouterKey = "issue"

type MsgIssueTransferOwnership struct {
	IssueId     string         `json:"issue_id" yaml:"issue_id"`
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address" yaml:"to_address"`
}

// Route Implements Msg.
func (msg MsgIssueTransferOwnership) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueTransferOwnership) Type() string { return "issue_transfer_ownership" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueTransferOwnership) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueTransferOwnership) GetSignBytes() []byte {
	return sdk.MustSortJSON(Cdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgIssueTransferOwnership) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}
