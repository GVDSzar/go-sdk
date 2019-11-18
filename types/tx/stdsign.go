package tx

import (
	"encoding/json"
	"github.com/tendermint/go-amino"
	"go-sdk/common/types"

	"github.com/tendermint/tendermint/crypto"
	"go-sdk/types/msg"
)

// StdSignDoc def
type StdSignDoc struct {
	ChainID       string            `json:"chain_id"`
	AccountNumber int64             `json:"account_number"`
	Sequence      int64             `json:"sequence"`
	Memo          string            `json:"memo"`
	Source        int64             `json:"source"`
	Msgs          []json.RawMessage `json:"msgs"`
	Data          []byte            `json:"data"`
}

// StdSignDoc def
type SignDocWithFee struct {
	ChainID       string            `json:"chain_id"`
	AccountNumber int64             `json:"account_number"`
	Sequence      int64             `json:"sequence"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Fee           msg.StdFee        `json:"fee"`
}

// StdSignMsg def
type StdSignMsg struct {
	ChainID       string     `json:"chain_id"`
	AccountNumber int64      `json:"account_number"`
	Sequence      int64      `json:"sequence"`
	Msgs          []msg.Msg  `json:"msgs"`
	Memo          string     `json:"memo"`
	Source        int64      `json:"source"`
	Data          []byte     `json:"data"`
	Fee           msg.StdFee `json:"fee"`
}

// StdSignature Signature
type StdSignature struct {
	crypto.PubKey `json:"pub_key"` // optional
	Signature     []byte `json:"signature"`
	AccountNumber int64  `json:"account_number"`
	Sequence      int64  `json:"sequence"`
}

//StdSignature structure (defined above) differs from original StdSignature (defined in cosmosSdk stdtx.go).
//So StdTxBasic structure represents original cosmosSDK structure
type StdSignatureBasic struct {
	crypto.PubKey `json:"pub_key"` // optional
	Signature     []byte `json:"signature"`
}

// Bytes gets message bytes
func (msg StdSignMsg) Bytes() []byte {
	return StdSignBytes(msg.ChainID, msg.AccountNumber, msg.Sequence, msg.Msgs, msg.Memo, msg.Source, msg.Data)
}

// StdSignBytes returns the bytes to sign for a transaction.
func StdSignBytes(chainID string, accnum int64, sequence int64, msgs []msg.Msg, memo string, source int64, data []byte) []byte {
	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}
	bz, err := Cdc.MarshalJSON(StdSignDoc{
		AccountNumber: accnum,
		ChainID:       chainID,
		Memo:          memo,
		Msgs:          msgsBytes,
		Sequence:      sequence,
		Source:        source,
		Data:          data,
	})
	if err != nil {
		panic(err)
	}
	return msg.MustSortJSON(bz)
}

// StdSignBytes returns the bytes to sign for a transaction.
func StdSignBytesWithFee(chainID string, accnum int64, sequence int64, msgs []msg.Msg, memo string, fee msg.StdFee) []byte {
	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}

	// workaround for a correct json body
	if fee.Amount == nil {
		fee.Amount = make(types.Coins, 0)
	}

	bz, err := amino.MarshalJSON(SignDocWithFee{
		AccountNumber: accnum,
		ChainID:       chainID,
		Memo:          memo,
		Msgs:          msgsBytes,
		Sequence:      sequence,
		Fee:           fee,
	})
	if err != nil {
		panic(err)
	}
	return msg.MustSortJSON(bz)
}
