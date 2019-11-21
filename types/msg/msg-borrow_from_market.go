package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type MsgBorrowFromMarket struct {
	Market           string           `json:"market"`
	BorrowTokens     types.Coins      `json:"borrowTokens"`
	CollateralTokens types.Coins      `json:"collateralTokens"`
	Supplier         types.AccAddress `json:"supplier"`
}

// Route should return the name of the module
func (msg MsgBorrowFromMarket) Route() string { return compoundRouterKey }

// Type should return the action
func (msg MsgBorrowFromMarket) Type() string { return "borrow_from_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBorrowFromMarket) ValidateBasic() error {
	if len(msg.Market) == 0 {
		return errors.New("Market cannot be empty")
	}
	if len(msg.BorrowTokens) == 0 {
		return errors.New("You must borrow at least one token")
	}
	if len(msg.CollateralTokens) == 0 {
		return errors.New("You must supply at least one collateral token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBorrowFromMarket) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners defines whose signature is required
func (msg MsgBorrowFromMarket) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Supplier}
}

func (msg MsgBorrowFromMarket) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Supplier}
}
