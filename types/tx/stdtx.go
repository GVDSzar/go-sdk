package tx

import (
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	"go-sdk/types/msg"
)

const Source int64 = 0

type Tx interface {
	// Gets the Msg.
	GetMsgs() []msg.Msg
}

// StdTx def
type StdTx struct {
	Msgs       []msg.Msg      `json:"msg"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
	Source     int64          `json:"source"`
	Data       []byte         `json:"data"`
}

//StdTx structure (defined above) differs from original StdTx (defined in cosmosSdk StdTx).
//So StdTxBasic structure represents original cosmosSDK structure
type StdTxBasic struct {
	Msgs       []msg.Msg           `json:"msg" yaml:"msg"`
	Fee        msg.StdFee          `json:"fee" yaml:"fee"`
	Signatures []StdSignatureBasic `json:"signatures" yaml:"signatures"`
	Memo       string              `json:"memo" yaml:"memo"`
}

// Represents TX structure for a default /txs broadcast endpoint
type TxBasic struct {
	Tx   StdTxBasic `json:"tx"`
	Mode string     `json:"mode"`
}

// NewStdTx to instantiate an instance
func NewStdTx(msgs []msg.Msg, sigs []StdSignature, memo string, source int64, data []byte) StdTx {
	return StdTx{
		Msgs:       msgs,
		Signatures: sigs,
		Memo:       memo,
		Source:     source,
		Data:       data,
	}
}

func NewStdxBasic(msgs []msg.Msg, sigs []StdSignatureBasic, memo string, fee msg.StdFee) StdTxBasic {
	return StdTxBasic{
		Msgs:       msgs,
		Signatures: sigs,
		Memo:       memo,
		Fee:        fee,
	}
}

func NewDefaultTransaction(msgs []msg.Msg, sigs []StdSignatureBasic, fee msg.StdFee, memo, mode string) *TxBasic {
	stdTx := StdTxBasic{
		Msgs:       msgs,
		Signatures: sigs,
		Memo:       memo,
		Fee:        fee,
	}
	return &TxBasic {
		stdTx,
		mode,
	}
}

// GetMsgs def
func (tx StdTx) GetMsgs() []msg.Msg { return tx.Msgs }

func (tx StdTxBasic) GetMsgs() []msg.Msg { return tx.Msgs }

type Info struct {
	Hash   common.HexBytes         `json:"hash"`
	Height int64                   `json:"height"`
	Tx     Tx                      `json:"tx"`
	Result types.ResponseDeliverTx `json:"result"`
}
