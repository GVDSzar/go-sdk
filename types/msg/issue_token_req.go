package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type DenominationsIssueTokenReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	SourceAddress string       `json:"source_address"`
	Owner         string       `json:"owner"`
	Name          string       `json:"name"`
	Symbol        string       `json:"symbol"`
	MaxSupply     int64        `json:"max_supply"`
	Mintable      bool         `json:"mintable"`
}

func NewDenominationsIssueTokenReq(br rest.BaseReq, sourceAddress, owner, name, symbol string, maxSupply int64, mintable bool) *DenominationsIssueTokenReq {
	return &DenominationsIssueTokenReq{
		br,
		sourceAddress,
		owner,
		name,
		symbol,
		maxSupply,
		mintable,
	}
}

func (d *DenominationsIssueTokenReq) ValidateMsg(res Msg) error {
	return nil
}

func (d DenominationsIssueTokenReq) Route() string { return DenominationsKey }

func (d DenominationsIssueTokenReq) Type() string { return DenominationsIssueTokenReqKey }

func (d DenominationsIssueTokenReq) ValidateBasic() error {
	return nil
}

func (d DenominationsIssueTokenReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d DenominationsIssueTokenReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (d DenominationsIssueTokenReq) GetInvolvedAddresses() []types.AccAddress {
	return d.GetSigners()
}
