package msg

import (
	"go-sdk/common/types"
	"time"
)


type MsgAddLiquidity struct {
	Deposit       types.Coin       `json:"deposit"`        // coin to be deposited as liquidity with an upper bound for its amount
	DepositAmount types.Int        `json:"deposit_amount"` // exact amount of native asset being add to the liquidity pool
	MinReward     types.Int        `json:"min_reward"`     // lower bound UNI sender is willing to accept for deposited coins
	Deadline      time.Time      `json:"deadline"`
	Sender        types.AccAddress `json:"sender"`
}

// NewMsgAddLiquidity creates a new MsgAddLiquidity object.
func NewMsgAddLiquidity(
	deposit types.Coin, depositAmount, minReward types.Int,
	deadline time.Time, sender types.AccAddress,
) MsgAddLiquidity {

	return MsgAddLiquidity{
		Deposit:       deposit,
		DepositAmount: depositAmount,
		MinReward:     minReward,
		Deadline:      deadline,
		Sender:        sender,
	}
}

// Type Implements Msg.
func (msg MsgAddLiquidity) Route() string { return uniswapRouterKey }

// Type Implements Msg.
func (msg MsgAddLiquidity) Type() string { return "add_liquidity" }

// ValidateBasic Implements Msg.
func (msg MsgAddLiquidity) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgAddLiquidity) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgAddLiquidity) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgAddLiquidity) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
