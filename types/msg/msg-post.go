package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go-sdk/common/types"
)

type MsgPost struct {
	Owner       types.AccAddress
	MarketID    sdk.Uint
	Direction   uint8
	Price       sdk.Uint
	Quantity    sdk.Uint
	TimeInForce uint16
}

