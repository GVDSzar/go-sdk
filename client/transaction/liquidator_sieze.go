package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) LiquidatorSieze(r *msg.SeizeAndStartCollateralAuctionRequest) (*tx.TxCommitResult, error) {

	return c.RestTransaction(r, liquidatorSeizeEndpoint, http.MethodPost)
}
