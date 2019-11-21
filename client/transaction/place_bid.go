package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) PlaceBidTx(r *msg.PlaceBidReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, placeBidEndpoint(r.AuctionID, r.Bidder, r.Bid, r.Lot), http.MethodPut)
}
