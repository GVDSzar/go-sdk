package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type DenominationsFreezeReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Amount  int64        `json:"amount"`
	Symbol  string       `json:"symbol"`
	Owner   string       `json:"owner"`
	Address string       `json:"address"`
}

func NewDenominationsFreezeReq(br rest.BaseReq, amount int64, symbol, owner string, address string) *DenominationsFreezeReq {
	return &DenominationsFreezeReq{
		br,
		amount,
		symbol,
		owner,
		address,
	}
}

func (d *DenominationsFreezeReq) ValidateMsg(res Msg) error {
	return nil
}

func (d DenominationsFreezeReq) Route() string { return DenominationsKey }

func (d DenominationsFreezeReq) Type() string { return DenominationsFreezeReqKey }

func (d DenominationsFreezeReq) ValidateBasic() error {
	return nil
}

func (d DenominationsFreezeReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d DenominationsFreezeReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d DenominationsFreezeReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
