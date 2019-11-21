package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type MsgRepayToMarket struct {
	Market      string           `json:"market"`
	RepayTokens types.Coins      `json:"repayTokens"`
	Borrower    types.AccAddress `json:"supplier"`
}

// Route should return the name of the module
func (msg MsgRepayToMarket) Route() string { return compoundRouterKey }

// Type should return the action
func (msg MsgRepayToMarket) Type() string { return "repay_to_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRepayToMarket) ValidateBasic() error {
	if len(msg.Market) == 0 {
		return errors.New("Market cannot be empty")
	}
	if len(msg.RepayTokens) == 0 {
		return errors.New("You must supply at least one token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRepayToMarket) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners defines whose signature is required
func (msg MsgRepayToMarket) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Borrower}
}

func (msg MsgRepayToMarket) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Borrower}
}
