package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type DenominationsBurnReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Amount  int64        `json:"amount"`
	Symbol  string       `json:"symbol"`
	Owner   string       `json:"owner"`
}

func NewDenominationsBurnReq(br rest.BaseReq, amount int64, symbol, owner string) *DenominationsBurnReq {
	return &DenominationsBurnReq{
		br,
		amount,
		symbol,
		owner,
	}
}

func (d *DenominationsBurnReq) ValidateMsg(res Msg) error {
	return nil
}

func (d DenominationsBurnReq) Route() string { return DenominationsKey }

func (d DenominationsBurnReq) Type() string { return DenominationsBurnReqKey }

func (d DenominationsBurnReq) ValidateBasic() error {
	return nil
}

func (d DenominationsBurnReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d DenominationsBurnReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d DenominationsBurnReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
