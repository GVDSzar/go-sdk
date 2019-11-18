package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgSwap struct {
	Sender      sdk.AccAddress `json:"sender"`
	Asset       sdk.Coin       `json:"asset"`
	TargetDenom string         `json:"target_denom"`
}


type MsgSwapWithHeight struct {
	Sender      sdk.AccAddress `json:"sender"`
	Asset       sdk.Coin       `json:"asset"`
	TargetDenom string         `json:"target_denom"`
}
