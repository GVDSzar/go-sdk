package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) NftsBurn(r *msg.BurnNFTReq, denom, id string) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, BurnNftEndpoint(denom, id), http.MethodPut)
}
