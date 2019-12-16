package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

const nftsBurnKey = "nfts_burn"

type BurnNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Denom   string       `json:"denom"`
	ID      string       `json:"id"`
}

func NewBurnNFTReq(br rest.BaseReq, denom, id string) *BurnNFTReq {
	return &BurnNFTReq{
		br,
		denom,
		id,
	}
}

func (d *BurnNFTReq) ValidateMsg(res Msg) error {
	return nil
}

func (d BurnNFTReq) Route() string { return nftsKey }

func (d BurnNFTReq) Type() string { return nftsBurnKey }

func (d BurnNFTReq) ValidateBasic() error {
	return nil
}

func (d BurnNFTReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d BurnNFTReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d BurnNFTReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
