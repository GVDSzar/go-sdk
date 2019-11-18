package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CoinIssueInfo struct {
	IssueId            string         `json:"issue_id"`
	Issuer             sdk.AccAddress `json:"issuer"`
	Owner              sdk.AccAddress `json:"owner"`
	IssueTime          int64          `json:"issue_time"`
	Name               string         `json:"name"`
	Symbol             string         `json:"symbol"`
	TotalSupply        sdk.Int        `json:"total_supply"`
	Description        string         `json:"description"`
	BurnOwnerDisabled  bool           `json:"burn_owner_disabled"`
	BurnHolderDisabled bool           `json:"burn_holder_disabled"`
	BurnFromDisabled   bool           `json:"burn_from_disabled"`
	FreezeDisabled     bool           `json:"freeze_disabled"`
	MintingFinished    bool           `json:"minting_finished"`
}

type CoinIssues []CoinIssueInfo

type CoinIssueInfoWithHeight struct {
	Height int64      `json:"height"`
	Result CoinIssues `json:"result"`
}

