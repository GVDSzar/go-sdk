package msg

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
	"go-sdk/common/types"
)

type StdSignature struct {
	crypto.PubKey `json:"pub_key" yaml:"pub_key"` // optional
	Signature     []byte `json:"signature" yaml:"signature"`
}

type StdTx struct {
	Msgs       []Msg          `json:"msg" yaml:"msg"`
	Fee        auth.StdFee    `json:"fee" yaml:"fee"`
	Signatures []StdSignature `json:"signatures" yaml:"signatures"`
	Memo       string         `json:"memo" yaml:"memo"`
}

type StdFee struct {
	Amount types.Coins `json:"amount" yaml:"amount"`
	Gas    uint64      `json:"gas" yaml:"gas"`
}
