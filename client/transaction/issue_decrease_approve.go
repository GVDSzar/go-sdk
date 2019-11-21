package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueDecreaseApprovalTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueDecreaseApproval(issueId, accAddress, amount), http.MethodPost)
}
