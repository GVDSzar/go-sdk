package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CSDT struct {
	//ID             []byte                                    // removing IDs for now to make things simpler
	Owner            sdk.AccAddress `json:"owner"`             // Account that authorizes changes to the CSDT
	CollateralDenom  string         `json:"collateral_denom"`  // Type of collateral stored in this CSDT
	CollateralAmount string         `json:"collateral_amount"` // Amount of collateral stored in this CSDT
	Debt             string         `json:"debt"`              // Amount of stable coin drawn from this CSDT
}

func NewCSDT(owner sdk.AccAddress, collateralDenom string, collateralAmount, debt string) CSDT {
	return CSDT{owner, collateralDenom, collateralAmount, debt}
}

type CSDTs []CSDT

type CSDTsWithHeight struct {
	Height int64 `json:"height"`
	Result CSDTs `json:"result"`
}

type CsdtModuleParams struct {
	GlobalDebtLimit  string
	CollateralParams []CollateralParams
}

type CollateralParams struct {
	Denom            string // Coin name of collateral type
	LiquidationRatio string // The ratio (Collateral (priced in stable coin) / Debt) under which a CSDT will be liquidated
	DebtLimit        string // Maximum amount of debt allowed to be drawn from this collateral type
	//DebtFloor        sdk.Int // used to prevent dust
}

type CSDTParamsWithHeight struct {
	Height int64            `json:"height"`
	Result CsdtModuleParams `json:"result"`
}
