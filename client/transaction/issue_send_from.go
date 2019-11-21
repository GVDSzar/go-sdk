package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueSendFrom(r *msg.PostIssueReq, issueId, from, to, amount string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueSendFrom(issueId, from, to, amount), http.MethodPost)
}
