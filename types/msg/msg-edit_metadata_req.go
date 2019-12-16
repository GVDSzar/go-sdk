package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type EditNFTMetadataReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Denom    string       `json:"denom"`
	ID       string       `json:"id"`
	TokenURI string       `json:"tokenURI"`
}

const nftsEditMetadataKey = "nfts_edit"

func NewEditNFTMetadataReq(br rest.BaseReq, denom, id, tokenUri string) *EditNFTMetadataReq {
	return &EditNFTMetadataReq{
		br,
		denom,
		id,
		tokenUri,
	}
}

func (d *EditNFTMetadataReq) ValidateMsg(res Msg) error {
	return nil
}

func (d EditNFTMetadataReq) Route() string { return nftsKey }

func (d EditNFTMetadataReq) Type() string { return nftsEditMetadataKey }

func (d EditNFTMetadataReq) ValidateBasic() error {
	return nil
}

func (d EditNFTMetadataReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d EditNFTMetadataReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d EditNFTMetadataReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
