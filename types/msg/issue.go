package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

const (
	CoinNameMinLength   = 3
	CoinNameMaxLength   = 32
	CoinSymbolMinLength = 2
	CoinSymbolMaxLength = 8
)

type MsgIssue struct {
	FromAddress  types.AccAddress `json:"from_address" yaml:"from_address"`
	*IssueParams `json:"params" yaml:"params"`
}

// Route Implements Msg.
func (msg MsgIssue) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssue) Type() string { return "issue" }

// ValidateBasic ensures addresses are valid and Coin is positive
func (msg MsgIssue) ValidateBasic() error {
	if len(msg.FromAddress) == 0 {
		return errors.New("missing sender address")
	}
	// Cannot issue zero or negative coins
	if msg.TotalSupply.IsZero() || !msg.TotalSupply.IsPositive() {
		return errors.New("cannot issue 0 or less supply")
	}
	if msg.TotalSupply.GT(CoinMaxTotalSupply) {
		return errors.New("total supply max value not valid")
	}
	if len(msg.Name) < CoinNameMinLength || len(msg.Name) > CoinNameMaxLength {
		return errors.New("coin name not valid")
	}
	if len(msg.Symbol) < CoinSymbolMinLength || len(msg.Symbol) > CoinSymbolMaxLength {
		return errors.New("coin symbol not valid")
	}
	if len(msg.Description) > CoinDescriptionMaxLength {
		return errors.New("coin description not valid")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssue) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssue) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

func (msg MsgIssue) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
