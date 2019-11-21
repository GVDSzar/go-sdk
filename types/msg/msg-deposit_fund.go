package msg

import "go-sdk/common/types"

type MsgDepositFund struct {
	Sender types.AccAddress
	Amount types.Coin
}

func NewMsgDepositFund(sender types.AccAddress, amount types.Coin) MsgDepositFund {
	return MsgDepositFund{
		Sender: sender,
		Amount: amount,
	}
}

// Route return the message type used for routing the message.
func (msg MsgDepositFund) Route() string { return "pool" }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgDepositFund) Type() string { return "deposit_fund" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgDepositFund) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgDepositFund) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgDepositFund) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgDepositFund) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
