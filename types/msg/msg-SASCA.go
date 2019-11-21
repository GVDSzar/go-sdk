package msg

import (
	"go-sdk/common/types"
)

type MsgSeizeAndStartCollateralAuction struct {
	Sender          types.AccAddress // only needed to pay the tx fees
	CsdtOwner       types.AccAddress
	CollateralDenom string
}

// Route return the message type used for routing the message.
func (msg MsgSeizeAndStartCollateralAuction) Route() string { return "liquidator" }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgSeizeAndStartCollateralAuction) Type() string { return "seize_and_start_auction" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgSeizeAndStartCollateralAuction) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgSeizeAndStartCollateralAuction) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgSeizeAndStartCollateralAuction) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
