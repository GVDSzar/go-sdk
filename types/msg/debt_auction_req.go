package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type StartDebtAuctionRequest struct {
	BaseReq rest.BaseReq     `json:"base_req"`
	Sender  types.AccAddress `json:"sender"`
}

func NewDebtAuctionRequest(br rest.BaseReq, sender types.AccAddress) *StartDebtAuctionRequest {
	return &StartDebtAuctionRequest {br, sender}
}

// Route should return the name of the module
func (msg StartDebtAuctionRequest) Route() string { return "nameservice" }

// Type should return the action
func (msg StartDebtAuctionRequest) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg StartDebtAuctionRequest) ValidateBasic() error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg StartDebtAuctionRequest) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// GetSigners defines whose signature is required
func (msg StartDebtAuctionRequest) GetSigners() []types.AccAddress {
	return nil
}

func (msg StartDebtAuctionRequest) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}

func (m StartDebtAuctionRequest) ValidateMsg(res Msg) error {
	// validate incoming msg
	return nil
}
