package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) DenominationsFreeze(r *msg.DenominationsFreezeReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, denominationsFreezeEndpoint, http.MethodPut)
}
