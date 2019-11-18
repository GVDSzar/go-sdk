package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) PriceRequest(r *msg.PriceReq, issueId string) (*tx.TxCommitResult, error) {

	return c.RestTransaction(r, issueDescription(issueId), http.MethodPost)
}
