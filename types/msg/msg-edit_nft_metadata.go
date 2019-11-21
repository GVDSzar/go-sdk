package msg

import (
	"go-sdk/common/types"
)

const nftRouterKey = "nft"

type MsgEditNFTMetadata struct {
	Sender   types.AccAddress
	ID       string
	Denom    string
	TokenURI string
}

// Route Implements Msg
func (msg MsgEditNFTMetadata) Route() string { return nftRouterKey }

// Type Implements Msg
func (msg MsgEditNFTMetadata) Type() string { return "edit_nft_metadata" }

// ValidateBasic Implements Msg.
func (msg MsgEditNFTMetadata) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgEditNFTMetadata) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgEditNFTMetadata) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgEditNFTMetadata) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
