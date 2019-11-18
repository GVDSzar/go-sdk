package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueIncreaseApprovalTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxCommitResult, error) {

	return c.RestTransaction(r, issueIncreaseApproval(issueId, accAddress, amount), http.MethodPost)
}
