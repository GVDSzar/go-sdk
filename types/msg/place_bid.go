package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go-sdk/common/types"
)

type PlaceBid struct {
	AuctionID uint64
	Bidder    sdk.AccAddress
	Bid       sdk.Coin
	Lot       sdk.Coin
}

func NewPlaceBid(auctionID uint64, bidder sdk.AccAddress, bid, lot sdk.Coin) *PlaceBid {
	return &PlaceBid{
		auctionID,
		bidder,
		bid,
		lot,
	}
}

func (r PlaceBid) Route() string { return "nameservice" }

func (r PlaceBid) Type() string { return "place_bid_req" }

func (r PlaceBid) ValidateBasic() error {
	// add validaton if needed
	return nil
}

func (r PlaceBid) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(r)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg PlaceBid) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (msg PlaceBid) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
