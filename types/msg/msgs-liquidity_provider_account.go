package msg

import (
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"go-sdk/common/types"
)

type LiquidityProviderAccount struct {
	exported.Account

	Credit types.Coins `json:"credit" yaml:"credit"`
}

