package msg

import "go-sdk/common/types"

type MsgWithdrawFund struct {
	Sender types.AccAddress
	Amount types.Coin
}

func NewMsgWithdrawFund(sender types.AccAddress, amount types.Coin) MsgWithdrawFund {
	return MsgWithdrawFund{
		Sender: sender,
		Amount: amount,
	}
}

// Route return the message type used for routing the message.
func (msg MsgWithdrawFund) Route() string { return "pool" }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgWithdrawFund) Type() string { return "withdraw_fund" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgWithdrawFund) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgWithdrawFund) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdrawFund) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgWithdrawFund) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
