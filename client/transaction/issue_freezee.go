package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueFreeze(r *msg.PostIssueReq, freezeType, issueId, accAddress string) (*tx.TxCommitResult, error) {

	return c.RestTransaction(r, issueFreeze(freezeType, issueId, accAddress), http.MethodPost)
}
