package msg

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type PlaceBidReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	AuctionID string       `json:"auction_id"`
	Bidder    string       `json:"bidder"`
	Bid       string       `json:"bid"`
	Lot       string       `json:"lot"`
}

func NewPlaceBidReq(br rest.BaseReq, auctionID, bidder, bid, lot string) *PlaceBidReq {
	return &PlaceBidReq{
		br,
		auctionID,
		bidder,
		bid,
		lot,
	}
}

func (m *PlaceBidReq) ValidateMsg(res Msg) error {
	// validate incoming msg
	return nil
}

func (rtx PlaceBidReq) ValidateResponse(b []byte) (Msg, error) {
	var bnr = new(BuyName)

	err := json.Unmarshal(b, bnr)
	if err != nil {
		return nil, err
	}

	return bnr, err
}

func (r PlaceBidReq) Route() string { return "auction" }

func (r PlaceBidReq) Type() string { return "place_bid" }

func (r PlaceBidReq) ValidateBasic() error {
	// add validaton if needed
	return nil
}

func (r PlaceBidReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(r)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg PlaceBidReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (msg PlaceBidReq) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
