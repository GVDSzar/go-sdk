package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) DenominationsIssueToken(r *msg.DenominationsIssueTokenReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, denominationsIssueToken, http.MethodPost)
}
