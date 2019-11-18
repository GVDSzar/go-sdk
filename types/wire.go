package types

import (
	"github.com/tendermint/go-amino"
	types "github.com/tendermint/tendermint/rpc/core/types"
	ntypes "go-sdk/common/types"
	"go-sdk/types/tx"
)

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	types.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}
