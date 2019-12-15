package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type DenominationsUnfreezeReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Amount  int64        `json:"amount"`
	Symbol  string       `json:"symbol"`
	Owner   string       `json:"owner"`
	Address string       `json:"address"`
}

func NewDenominationsUnfreezeReq(br rest.BaseReq, amount int64, symbol, owner string, address string) *DenominationsUnfreezeReq {
	return &DenominationsUnfreezeReq{
		br,
		amount,
		symbol,
		owner,
		address,
	}
}

func (d *DenominationsUnfreezeReq) ValidateMsg(res Msg) error {
	return nil
}

func (d DenominationsUnfreezeReq) Route() string { return DenominationsKey }

func (d DenominationsUnfreezeReq) Type() string { return DenominationsUnFreezeReqKey }

func (d DenominationsUnfreezeReq) ValidateBasic() error {
	return nil
}

func (d DenominationsUnfreezeReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d DenominationsUnfreezeReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d DenominationsUnfreezeReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
