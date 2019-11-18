package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CurrentPrice struct {
	AssetCode string  `json:"asset_code"`
	Price     sdk.Dec `json:"price"`
	Expiry    sdk.Int `json:"expiry"`
}

type CurrentPriceWithHeight struct {
	Height int64          `json:"height"`
	Result []CurrentPrice `json:"result"`
}
