package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueBurnTx(r *msg.PostIssueReq, issueId, amount string) (*tx.TxCommitResult, error) {

	return c.RestTransaction(r, issueBurn(issueId, amount), http.MethodPost)
}
