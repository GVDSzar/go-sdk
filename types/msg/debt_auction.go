package msg

import (
	"go-sdk/common/types"
)

type MsgStartDebtAuction struct {
	Sender  string `json:"Sender"`
}

func NewDebtAuction(sender types.AccAddress) *MsgStartDebtAuction {
	return &MsgStartDebtAuction{sender.String()}
}

// Route should return the name of the module
func (msg MsgStartDebtAuction) Route() string { return "nameservice" }

// Type should return the action
func (msg MsgStartDebtAuction) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgStartDebtAuction) ValidateBasic() error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgStartDebtAuction) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// GetSigners defines whose signature is required
func (msg MsgStartDebtAuction) GetSigners() []types.AccAddress {
	return nil
}

func (msg MsgStartDebtAuction) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}