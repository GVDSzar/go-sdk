package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueDisableFeature(r *msg.PostIssueReq, issueId, feature string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueDisableFeature(issueId, feature), http.MethodPost)
}
