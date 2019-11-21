package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

const compoundRouterKey = "compound"

type MsgCreateCompound struct {
	Name            string           `json:"name"`
	Denom           string           `json:"denom"`
	InterestRate    types.Coins      `json:"interest_tate"`
	Buyer           types.AccAddress `json:"buyer"`
	TokenName       string           `json:"token_name"`
	CollateralToken string           `json:"collateral_token"`
}

// Route should return the name of the module
func (msg MsgCreateCompound) Route() string { return compoundRouterKey }

// Type should return the action
func (msg MsgCreateCompound) Type() string { return "create_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateCompound) ValidateBasic() error {
	if len(msg.Buyer) == 0 {
		return errors.New("invalid address")
	}
	if len(msg.Name) == 0 {
		return errors.New("Name cannot be empty")
	}
	if len(msg.Denom) == 0 {
		return errors.New("Symbol cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateCompound) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners defines whose signature is required
func (msg MsgCreateCompound) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Buyer}
}

func (msg MsgCreateCompound) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Buyer}
}
