package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) DenominationsBurn(r *msg.DenominationsBurnReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, denominationsBurnEndpoint, http.MethodPut)
}
