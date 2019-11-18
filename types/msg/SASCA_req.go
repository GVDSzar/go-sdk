package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

// structure name is taken from original Xar project, but it appears to be too complicated to be read easily
// for the purpose of simplification another name will be considered
type SeizeAndStartCollateralAuctionRequest struct {
	BaseReq         rest.BaseReq     `json:"base_req"`
	Sender          types.AccAddress `json:"sender"`
	CsdtOwner       types.AccAddress `json:"csdt_owner"`
	CollateralDenom string           `json:"collateral_denom"`
}

func NewSASCARequest(baseReq rest.BaseReq, sender types.AccAddress, csdtOwner types.AccAddress, collateralDenom string) *SeizeAndStartCollateralAuctionRequest {
	return &SeizeAndStartCollateralAuctionRequest{
		baseReq, sender, csdtOwner, collateralDenom,
	}
}

func (r SeizeAndStartCollateralAuctionRequest) Route() string { return "undefined" }

func (r SeizeAndStartCollateralAuctionRequest) Type() string { return "undefined" }

func (r SeizeAndStartCollateralAuctionRequest) ValidateBasic() error {
	// add validaton if needed
	return nil
}

func (r SeizeAndStartCollateralAuctionRequest) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(r)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg SeizeAndStartCollateralAuctionRequest) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (msg SeizeAndStartCollateralAuctionRequest) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}

func (m *SeizeAndStartCollateralAuctionRequest) ValidateMsg(res Msg) error {
	// validate incoming msg
	return nil
}
