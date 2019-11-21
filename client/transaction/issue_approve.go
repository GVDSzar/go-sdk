package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueApproveTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxBroadcastResult, error) {

	return c.RestTransaction(r, issueApprove(issueId, accAddress, amount), http.MethodPost)
}
