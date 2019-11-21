package msg

import (
	"go-sdk/common/types"
)

type MsgCreateOrModifyCSDT struct {
	Sender           types.AccAddress
	CollateralDenom  string
	CollateralChange types.Int
	DebtChange       types.Int
}

// NewMsgPlaceBid returns a new MsgPlaceBid.
func NewMsgCreateOrModifyCSDT(sender types.AccAddress, collateralDenom string, collateralChange types.Int, debtChange types.Int) MsgCreateOrModifyCSDT {
	return MsgCreateOrModifyCSDT{
		Sender:           sender,
		CollateralDenom:  collateralDenom,
		CollateralChange: collateralChange,
		DebtChange:       debtChange,
	}
}

// Route return the message type used for routing the message.
func (msg MsgCreateOrModifyCSDT) Route() string { return "csdt" }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgCreateOrModifyCSDT) Type() string { return "create_modify_csdt" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgCreateOrModifyCSDT) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgCreateOrModifyCSDT) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCreateOrModifyCSDT) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgCreateOrModifyCSDT) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
