package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type DenominationsMintReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Amount  int64        `json:"amount"`
	Symbol  string       `json:"symbol"`
	Owner   string       `json:"owner"`
}

func NewDenominationsMintReq(br rest.BaseReq, amount int64, symbol, owner string) *DenominationsMintReq {
	return &DenominationsMintReq{
		br,
		amount,
		symbol,
		owner,
	}
}

func (d *DenominationsMintReq) ValidateMsg(res Msg) error {
	return nil
}

func (d DenominationsMintReq) Route() string { return DenominationsKey }

func (d DenominationsMintReq) Type() string { return DenominationsMintReqKey }

func (d DenominationsMintReq) ValidateBasic() error {
	return nil
}

func (d DenominationsMintReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d DenominationsMintReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d DenominationsMintReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
