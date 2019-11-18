package tx

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"

	"go-sdk/types/msg"
)

// cdc global variable
var Cdc = amino.NewCodec()

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
	cdc.RegisterConcrete(StdTxBasic{}, "cosmos-sdk/StdTx", nil)

	msg.RegisterCodec(cdc)
}

func init() {
	cryptoAmino.RegisterAmino(Cdc)
	RegisterCodec(Cdc)
}
