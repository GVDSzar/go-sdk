package msg

import (
	"go-sdk/common/types"
)

type (
	MsgMintTokens struct {
		Amount            types.Coins
		LiquidityProvider types.AccAddress
	}

	MsgBurnTokens struct {
		Amount            types.Coins
		LiquidityProvider types.AccAddress
	}
)

