package msg

import (
	"go-sdk/common/types"
	"time"
)


type MsgRemoveLiquidity struct {
	Withdraw       types.Coin       `json:"withdraw"`        // coin to be withdrawn with a lower bound for its amount
	WithdrawAmount types.Int        `json:"withdraw_amount"` // amount of UNI to be burned to withdraw liquidity from a reserve pool
	MinNative      types.Int        `json:"min_native"`      // minimum amount of the native asset the sender is willing to accept
	Deadline       time.Time      `json:"deadline"`
	Sender         types.AccAddress `json:"sender"`
}

// NewMsgRemoveLiquidity creates a new MsgRemoveLiquidity object
func NewMsgRemoveLiquidity(
	withdraw types.Coin, withdrawAmount, minNative types.Int,
	deadline time.Time, sender types.AccAddress,
) MsgRemoveLiquidity {

	return MsgRemoveLiquidity{
		Withdraw:       withdraw,
		WithdrawAmount: withdrawAmount,
		MinNative:      minNative,
		Deadline:       deadline,
		Sender:         sender,
	}
}

// Type Implements Msg.
func (msg MsgRemoveLiquidity) Route() string { return uniswapRouterKey }

// Type Implements Msg.
func (msg MsgRemoveLiquidity) Type() string { return "remove_liquidity" }

// ValidateBasic Implements Msg.
func (msg MsgRemoveLiquidity) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgRemoveLiquidity) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgRemoveLiquidity) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgRemoveLiquidity) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}


