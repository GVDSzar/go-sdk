package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go-sdk/common/types"
)

type MsgCancel struct {
	Owner   types.AccAddress
	OrderID sdk.Uint
}
