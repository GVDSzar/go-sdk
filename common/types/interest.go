package types

import (
	"time"
)

type Interest struct {
	LastAppliedTime   time.Time      `json:"last_applied" yaml:"last_applied"`
	LastAppliedHeight string        `json:"last_applied_height" yaml:"last_applied_height"`
	InterestAssets    InterestAssets `json:"assets" yaml:"assets"`
}

type InterestAsset struct {
	Denom    string  `json:"denom" yaml:"denom"`
	Interest string `json:"interest" yaml:"interest"`
	Accum    string `json:"accum" yaml:"accum"`
}

type InterestAssets []InterestAsset

type InterestWithHeight struct {
	Height int64    `json:"height"`
	Result Interest `json:"result"`
}
