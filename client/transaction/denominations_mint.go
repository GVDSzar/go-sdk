package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) DenominationsMint(r *msg.DenominationsMintReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, denominationsMintEndpoint, http.MethodPut)
}
