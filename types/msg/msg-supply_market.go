package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type MsgSupplyMarket struct {
	Market     string           `json:"compound"`
	LendTokens types.Coins      `json:"lend_tokens"`
	Supplier   types.AccAddress `json:"supplier"`
}

// NewMsgCreateMarket is the constructor function for MsgBuyName
func NewMsgSupplyMarket(market string, coins types.Coins, supplier types.AccAddress) MsgSupplyMarket {
	return MsgSupplyMarket{
		Market:     market,
		LendTokens: coins,
		Supplier:   supplier,
	}
}

// Route should return the name of the module
func (msg MsgSupplyMarket) Route() string { return compoundRouterKey }

// Type should return the action
func (msg MsgSupplyMarket) Type() string { return "supply_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSupplyMarket) ValidateBasic() error {
	if len(msg.Market) == 0 {
		return errors.New("Market cannot be empty")
	}
	if len(msg.LendTokens) == 0 {
		return errors.New("You must supply at least one token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSupplyMarket) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners defines whose signature is required
func (msg MsgSupplyMarket) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Supplier}
}

func (msg MsgSupplyMarket) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Supplier}
}
