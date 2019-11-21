package msg

import (
	"go-sdk/common/types"
)

type BaseNFT struct {
	ID       string           `json:"id,omitempty" yaml:"id"`     // id of the token; not exported to clients
	Owner    types.AccAddress `json:"owner" yaml:"owner"`         // account address that owns the NFT
	TokenURI string           `json:"token_uri" yaml:"token_uri"` // optional extra properties available for querying
}
