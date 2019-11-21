package msg

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

var CoinMaxTotalSupply, _ = types.NewIntFromString("1000000000000000000000000000000000000")

type MsgIssueMint struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Int        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueMint) Route() string { return "issue" }

// Type Implements Msg.
func (msg MsgIssueMint) Type() string { return "issue_mint" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueMint) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("issueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if !msg.Amount.IsPositive() {
		return errors.New("cannot mint 0 or negative coin amounts")
	}
	if msg.Amount.GT(CoinMaxTotalSupply) {
		return ErrorCoinTotalSupplyMaxValueNotValid
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueMint) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueMint) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}