package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

const nftsMintKey = "nfts_mint"

type MintNFTReq struct {
	BaseReq   rest.BaseReq   `json:"base_req"`
	Recipient types.AccAddress `json:"recipient"`
	Denom     string         `json:"denom"`
	ID        string         `json:"id"`
	TokenURI  string         `json:"tokenURI"`
}

func NewMintNFTReq(br rest.BaseReq, recipient types.AccAddress, denom, id, tokenUri string) *MintNFTReq {
	return &MintNFTReq{
		br,
		recipient,
		denom,
		id,
		tokenUri,
	}
}

func (d *MintNFTReq) ValidateMsg(res Msg) error {
	return nil
}

func (d MintNFTReq) Route() string { return nftsKey }

func (d MintNFTReq) Type() string { return nftsMintKey }

func (d MintNFTReq) ValidateBasic() error {
	return nil
}

func (d MintNFTReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d MintNFTReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d MintNFTReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
