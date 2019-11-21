package types

import (
	"github.com/tendermint/go-amino"
	types "github.com/tendermint/tendermint/rpc/core/types"
	ntypes "go-sdk/common/types"
	"go-sdk/types/tx"
	"sync"
)

var cdc *amino.Codec
var cdcMtx sync.Mutex

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	types.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}

func GetCodec() *amino.Codec {
	cdcMtx.Lock()
	if cdc == nil {
		cdc = NewCodec()
	}
	cdcMtx.Unlock()

	return cdc
}
