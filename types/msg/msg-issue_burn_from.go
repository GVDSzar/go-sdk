package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go-sdk/common/types"
)

type MsgIssueBurnFrom struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Int        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueBurnFrom) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnFrom) Type() string { return "issue_burn_from" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnFrom) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("issueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.Amount.I.Int64() < 0 {
		return sdk.ErrInvalidCoins("cannot Burn 0 or negative coin amounts")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurnFrom) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnFrom) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

func (msg MsgIssueBurnFrom) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
