package msg

import (
	"go-sdk/common/types"
)

type (
	MsgIncreaseCredit struct {
		CreditIncrease    types.Coins
		LiquidityProvider types.AccAddress
		Issuer            types.AccAddress
	}

	MsgDecreaseCredit struct {
		CreditDecrease    types.Coins
		LiquidityProvider types.AccAddress
		Issuer            types.AccAddress
	}

	MsgRevokeLiquidityProvider struct {
		LiquidityProvider types.AccAddress
		Issuer            types.AccAddress
	}

	MsgSetInterest struct {
		Denom        string
		InterestRate types.Dec
		Issuer       types.AccAddress
	}
)
