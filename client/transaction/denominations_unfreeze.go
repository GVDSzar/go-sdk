package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) DenominationsUnfreeze(r *msg.DenominationsUnfreezeReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, denominationsUnfreezeEndpoint, http.MethodPut)
}
