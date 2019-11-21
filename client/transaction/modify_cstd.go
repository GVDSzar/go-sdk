package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) ModifyCsdtTx(r *msg.ModifyCsdtReq) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, csdtParamsEndpoint, http.MethodPut)
}
