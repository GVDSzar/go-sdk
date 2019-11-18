package msg

import (
	"go-sdk/common/types"
)

type CreateOrModifyCSDT struct {
	Sender           types.AccAddress
	CollateralDenom  string
	CollateralChange types.Int
	DebtChange       types.Int
}

func NewCreateOrModifyCSDT(sender types.AccAddress, collateralDenom string, collateralChange types.Int, debtChange types.Int, ) *CreateOrModifyCSDT {
	return &CreateOrModifyCSDT{
		sender,
		collateralDenom,
		collateralChange,
		debtChange,
	}
}

func (m *CreateOrModifyCSDT) ValidateMsg(res Msg) error {
	// validate incoming msg
	return nil
}

func (r CreateOrModifyCSDT) Route() string { return "csdt" }

func (r CreateOrModifyCSDT) Type() string { return "create_modify_csdt" }

func (r CreateOrModifyCSDT) ValidateBasic() error {
	// add validaton if needed
	return nil
}

func (r CreateOrModifyCSDT) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(r)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg CreateOrModifyCSDT) GetSigners() []types.AccAddress {
	return nil
}

func (msg CreateOrModifyCSDT) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
