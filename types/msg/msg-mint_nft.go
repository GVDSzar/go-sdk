package msg

import (
	"go-sdk/common/types"
)

type MsgMintNFT struct {
	Sender    types.AccAddress
	Recipient types.AccAddress
	ID        string
	Denom     string
	TokenURI  string
}

// Route Implements Msg
func (msg MsgMintNFT) Route() string { return nftRouterKey }

// Type Implements Msg
func (msg MsgMintNFT) Type() string { return "mint_nft" }

// ValidateBasic Implements Msg.
func (msg MsgMintNFT) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMintNFT) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgMintNFT) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgMintNFT) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
