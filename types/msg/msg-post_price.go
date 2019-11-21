package msg

import (
	"go-sdk/common/types"
)

const oracleRouterKey = "oracle"

type MsgPostPrice struct {
	From      types.AccAddress // client that sent in this address
	AssetCode string         // asset code used by exchanges/api
	Price     types.Dec        // price in decimal (max precision 18)
	Expiry    types.Int        // block height
}

// Route Implements Msg.
func (msg MsgPostPrice) Route() string { return oracleRouterKey }

// Type Implements Msg
func (msg MsgPostPrice) Type() string { return "post_price" }

// GetSignBytes Implements Msg.
func (msg MsgPostPrice) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgPostPrice) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.From}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgPostPrice) ValidateBasic() error {
	return nil
}

func (msg MsgPostPrice) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.From}
}
