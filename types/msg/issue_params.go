package msg

import "github.com/cosmos/cosmos-sdk/types"

type IssueParams struct {
	Name               string    `json:"name"`
	Symbol             string    `json:"symbol"`
	TotalSupply        types.Int `json:"total_supply"`
	Description        string    `json:"description"`
	BurnOwnerDisabled  bool      `json:"burn_owner_disabled"`
	BurnHolderDisabled bool      `json:"burn_holder_disabled"`
	BurnFromDisabled   bool      `json:"burn_from_disabled"`
	MintingFinished    bool      `json:"minting_finished"`
	FreezeDisabled     bool      `json:"freeze_disabled"`
}

const (
	BurnOwnerMarker = 1 << iota
	BurnHolderMarker
	BurnFromMarker
	MintingMarker
	FreezeMarker
)

func GetIssueParamsBitMask(burnOwnerDisabled, burnHolderDisabled, burnFromDisabled, mintingFinished, freezeDisabled bool) uint8 {
	var mask uint8 = 0
	if burnOwnerDisabled {
		mask = mask | BurnOwnerMarker
	}
	if burnHolderDisabled {
		mask = mask | BurnHolderMarker
	}
	if burnFromDisabled {
		mask = mask | BurnFromMarker
	}
	if mintingFinished {
		mask = mask | MintingMarker
	}
	if freezeDisabled {
		mask = mask | FreezeMarker
	}
	return mask
}


// this function is made for code-style purposes only
// use it if shorter signature is needed
// BM here is a shortened "bitmask"
func NewIssueParamsBm(name, symbol, description string, totalSupply types.Int, bitMask uint8) *IssueParams {
	return &IssueParams{
		Name:               name,
		Symbol:             symbol,
		Description:        description,
		TotalSupply:        totalSupply,
		BurnOwnerDisabled:  bitMask&BurnOwnerMarker != 0,
		BurnHolderDisabled: bitMask&BurnHolderMarker != 0,
		BurnFromDisabled:   bitMask&BurnFromMarker != 0,
		MintingFinished:    bitMask&MintingMarker != 0,
		FreezeDisabled:     bitMask&FreezeMarker != 0,
	}
}

func NewIssueParams(name, symbol, description string, totalSupply types.Int, burnOwnerDisabled, burnHolderDisabled, burnFromDisabled, mintingFinished, freezeDisabled bool) *IssueParams {
	return &IssueParams{
		Name: name,
		Symbol: symbol,
		Description: description,
		TotalSupply: totalSupply,
		BurnOwnerDisabled: burnOwnerDisabled,
		BurnHolderDisabled: burnHolderDisabled,
		BurnFromDisabled: burnFromDisabled,
		MintingFinished: mintingFinished,
		FreezeDisabled: freezeDisabled,
	}
}