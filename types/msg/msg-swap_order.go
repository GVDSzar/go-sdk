package msg

import (
	"go-sdk/common/types"
	"time"
)

const uniswapRouterKey = "uniswap"

type MsgSwapOrder struct {
	Input      types.Coin       `json:"input"`    // the amount the sender is trading
	Output     types.Coin       `json:"output"`   // the amount the sender is recieivng
	Deadline   time.Time      `json:"deadline"` // deadline for the transaction to still be considered valid
	Sender     types.AccAddress `json:"sender"`
	Recipient  types.AccAddress `json:"recipient"`
	IsBuyOrder bool           `json:"is_buy_order"` // boolean indicating whether the order should be treated as a buy or sell
}

// NewMsgSwapOrder creates a new MsgSwapOrder object.
func NewMsgSwapOrder(
	input, output types.Coin, deadline time.Time,
	sender, recipient types.AccAddress, isBuyOrder bool,
) MsgSwapOrder {

	return MsgSwapOrder{
		Input:      input,
		Output:     output,
		Deadline:   deadline,
		Sender:     sender,
		Recipient:  recipient,
		IsBuyOrder: isBuyOrder,
	}
}

// Route Implements Msg.
func (msg MsgSwapOrder) Route() string { return uniswapRouterKey }

// Type Implements Msg.
func (msg MsgSwapOrder) Type() string { return "swap_order" }

// ValidateBasic Implements Msg.
func (msg MsgSwapOrder) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSwapOrder) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgSwapOrder) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgSwapOrder) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

