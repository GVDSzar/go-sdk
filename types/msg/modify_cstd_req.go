package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type ModifyCsdtReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Csdt    types.CSDT   `json:"csdt"`
}

func NewModifyCsdtReq(br rest.BaseReq, csdt types.CSDT) *ModifyCsdtReq {
	return &ModifyCsdtReq{
		br,
		csdt,
	}
}

func (m *ModifyCsdtReq) ValidateMsg(res Msg) error {
	// validate incoming msg
	return nil
}

func (r ModifyCsdtReq) Route() string { return "csdt" }

func (r ModifyCsdtReq) Type() string { return "create_modify_csdt" }

func (r ModifyCsdtReq) ValidateBasic() error {
	// add validaton if needed
	return nil
}

func (r ModifyCsdtReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(r)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg ModifyCsdtReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (msg ModifyCsdtReq) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
