package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type PriceReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	AssetCode string       `json:"asset_code"`
	Price     string       `json:"price"`
	Expiry    string       `json:"expiry"`
}

func NewPriceReq(br rest.BaseReq, assetCode string, price string, expiry string) *PriceReq {
	return &PriceReq{
		br,
		assetCode,
		price,
		expiry,
	}
}

func (r PriceReq) Route() string { return "undefined" }

func (r PriceReq) Type() string { return "undefined" }

func (r PriceReq) ValidateBasic() error {
	// add validaton if needed
	return nil
}

func (m *PriceReq) ValidateMsg(res Msg) error {
	// validate incoming msg
	return nil
}

func (r PriceReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(r)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg PriceReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (msg PriceReq) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
