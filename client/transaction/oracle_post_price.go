package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) OraclePostPrice(r *msg.DenominationsFreezeReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, PostPriceEndpoint(), http.MethodPut)
}
