package msg

import "go-sdk/common/types"

type ID uint64
type EndTime int64

type BaseAuction struct {
	ID         ID
	Initiator  types.AccAddress
	Lot        types.Coin
	Bidder     types.AccAddress
	Bid        types.Coin
	EndTime    EndTime
	MaxEndTime EndTime
}

type ForwardAuction struct {
	BaseAuction
}

type ReverseAuction struct {
	BaseAuction
}

type ForwardReverseAuction struct {
	BaseAuction
	MaxBid      types.Coin
	OtherPerson types.AccAddress
}
