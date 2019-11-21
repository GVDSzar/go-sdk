package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueBurnFromTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueBurnFrom(issueId, accAddress, amount), http.MethodPost)
}
