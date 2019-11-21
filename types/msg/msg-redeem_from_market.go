package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type MsgRedeemFromMarket struct {
	Market       string           `json:"market"`
	RedeemTokens types.Coins      `json:"redeemTokens"`
	Supplier     types.AccAddress `json:"supplier"`
}

// Route should return the name of the module
func (msg MsgRedeemFromMarket) Route() string { return compoundRouterKey }

// Type should return the action
func (msg MsgRedeemFromMarket) Type() string { return "redeem_from_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRedeemFromMarket) ValidateBasic() error {
	if len(msg.Market) == 0 {
		return errors.New("Market cannot be empty")
	}
	if len(msg.RedeemTokens) == 0 {
		return errors.New("You must supply at least one token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRedeemFromMarket) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners defines whose signature is required
func (msg MsgRedeemFromMarket) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Supplier}
}

func (msg MsgRedeemFromMarket) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Supplier}
}
