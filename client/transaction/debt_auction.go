package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) DebtAuction(r *msg.StartDebtAuctionRequest) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, liquidatorMintEndpoint, http.MethodPost)
}
