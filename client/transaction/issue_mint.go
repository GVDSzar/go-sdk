package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueMint(r *msg.PostIssueReq, issueId, amount, to string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueMint(issueId, amount, to), http.MethodPost)
}
