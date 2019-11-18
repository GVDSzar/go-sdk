package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Params struct {
	IssueFee         sdk.Coin `json:"issue_fee"`
	MintFee          sdk.Coin `json:"mint_fee"`
	FreezeFee        sdk.Coin `json:"freeze_fee"`
	UnFreezeFee      sdk.Coin `json:"unfreeze_fee"`
	BurnFee          sdk.Coin `json:"burn_fee"`
	BurnFromFee      sdk.Coin `json:"burn_from_fee"`
	TransferOwnerFee sdk.Coin `json:"transfer_owner_fee"`
	DescribeFee      sdk.Coin `json:"describe_fee"`
}

type ParamsWithHeight struct {
	Height int64    `json:"height"`
	Result []Params `json:"result"`
}
