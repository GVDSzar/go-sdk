package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueTransfer(r *msg.PostIssueReq, issueId, to string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueTransfer(issueId, to), http.MethodPost)
}
