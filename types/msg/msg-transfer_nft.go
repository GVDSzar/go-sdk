package msg

import "go-sdk/common/types"

type MsgTransferNFT struct {
	Sender    types.AccAddress
	Recipient types.AccAddress
	Denom     string
	ID        string
}
