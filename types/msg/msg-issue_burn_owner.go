package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgIssueBurnOwner struct {
	IssueId     string         `json:"issue_id" yaml:"issue_id"`
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Amount      sdk.Int        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueBurnOwner) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnOwner) Type() string { return "issue_burn_owner" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnOwner) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("issueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("cannot Burn 0 or negative coin amounts")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurnOwner) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}
