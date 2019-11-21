package msg

import (
	"go-sdk/common/types"
)

type MsgBurnNFT struct {
	Sender types.AccAddress
	ID     string
	Denom  string
}

// Route Implements Msg
func (msg MsgBurnNFT) Route() string { return nftRouterKey }

// Type Implements Msg
func (msg MsgBurnNFT) Type() string { return "burn_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBurnNFT) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBurnNFT) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgBurnNFT) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgBurnNFT) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
