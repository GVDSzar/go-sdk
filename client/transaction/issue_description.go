package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueDescription(r *msg.PostIssueReq, issueId string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueDescription(issueId), http.MethodPost)
}
