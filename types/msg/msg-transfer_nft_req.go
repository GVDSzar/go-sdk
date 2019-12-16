package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

const nftsKey = "nfts"
const nftsTransferKey = "nfts_transfer"

type TransferNFTReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Denom     string       `json:"denom"`
	ID        string       `json:"id"`
	Recipient string       `json:"recipient"`
}

func NewTransferNFTReq(br rest.BaseReq, denom, id, recipient string) *TransferNFTReq {
	return &TransferNFTReq{
		br,
		denom,
		id,
		recipient,
	}
}

func (d *TransferNFTReq) ValidateMsg(res Msg) error {
	return nil
}

func (d TransferNFTReq) Route() string { return nftsKey }

func (d TransferNFTReq) Type() string { return nftsTransferKey }

func (d TransferNFTReq) ValidateBasic() error {
	return nil
}

func (d TransferNFTReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d TransferNFTReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d TransferNFTReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
